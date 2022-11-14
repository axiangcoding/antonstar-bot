package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/service/crawler"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
	"gorm.io/gorm"
	"time"
)

type CrawlerResult struct {
	StartCrawlerSuccess bool   `json:"start_crawler_success"`
	ResponseStatus      int    `json:"response_status"`
	Found               bool   `json:"found"`
	Nick                string `json:"nick"`
	Data                any    `json:"data"`
}

type ScheduleForm struct {
	Nick     string                  `json:"nick,omitempty"`
	SendForm cqhttp.SendGroupMsgForm `json:"send_form"`
}

type ScheduleResult struct {
	NodeName string `json:"node_name,omitempty"`
	Status   string `json:"status,omitempty"`
	JobId    string `json:"jobid,omitempty"`
}

func WaitForCrawlerFinished(missionId string, fullMsg bool) error {
	totalDelay := 60
	duration := 3
	i := 0
	var detailForm ScheduleForm
	for i <= totalDelay {
		time.Sleep(time.Second * time.Duration(duration))
		mission, err := FindMission(missionId)
		if err != nil {
			logging.Warnf("polling find mission failed. %s", err)
			i += duration
			continue
		}
		_ = json.Unmarshal([]byte(mission.Detail), &detailForm)

		if mission.Status == table.MissionStatusSuccess {
			user, err := FindGameProfile(detailForm.Nick)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					detailForm.SendForm.Message = "未找到该用户，请检查游戏昵称是否正确"
					break
				}
			} else {
				if fullMsg {
					detailForm.SendForm.Message = user.ToDisplayGameUser().ToFriendlyFullString()
				} else {
					detailForm.SendForm.Message = user.ToDisplayGameUser().ToFriendlyShortString()
				}
				break
			}
		} else if mission.Status == table.MissionStatusFailed {
			detailForm.SendForm.Message = "查询失败，请稍后重试"
			break
		}
		i += duration
	}
	if i > totalDelay {
		detailForm.SendForm.Message = "对不起，查询超时，请稍后重试"
	}
	MustSendGroupMsg(detailForm.SendForm)
	return nil
}

func GetUserInfoFromWarThunder(missionId string, nick string) error {
	urlTemplate := "https://warthunder.com/zh/community/userinfo/?nick=%s"
	url := fmt.Sprintf(urlTemplate, nick)

	c := colly.NewCollector(
		colly.AllowedDomains("warthunder.com"),
		colly.MaxDepth(1),
		colly.IgnoreRobotsTxt(),
	)
	extensions.RandomUserAgent(c)

	c.OnHTML("div[class=user__unavailable-title]", func(e *colly.HTMLElement) {
		logging.Warnf("%s userinfo not found", nick)
		MustPutRefreshFlag(nick)
		MustFinishMissionWithResult(missionId, table.MissionStatusSuccess, CrawlerResult{
			Found: false,
			Nick:  nick,
		})
	})

	c.OnHTML("div[class=user-info]", func(e *colly.HTMLElement) {
		data := crawler.ExtractGaijinData(e)
		_, err := FindGameProfile(nick)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := SaveGameProfile(data); err != nil {
					logging.Warn(err)
				}
			} else {
				logging.Warn(err)
			}
		} else {
			if err := UpdateGameProfile(nick, data); err != nil {
				logging.Warn(err)
			}
		}

		// if err := GetUserInfoFromThunderskill(nick); err != nil {
		// 	logging.Warn("failed on update thunder skill profile. ", err)
		// }
		MustPutRefreshFlag(nick)
		MustFinishMissionWithResult(missionId, table.MissionStatusSuccess, CrawlerResult{
			Found: true,
			Nick:  nick,
			Data:  data},
		)
	})

	c.OnRequest(func(r *colly.Request) {
		logging.Infof("start visiting %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		logging.Warnf("visiting %s failed. %s", r.Request.URL.String(), err)
		MustFinishMissionWithResult(missionId, table.MissionStatusFailed, CrawlerResult{
			Found:          false,
			Nick:           nick,
			ResponseStatus: r.StatusCode,
		})
	})

	err := c.Post(url, nil)
	if err != nil {
		logging.Warn("colly post failed", err)
		MustFinishMissionWithResult(missionId, table.MissionStatusFailed, CrawlerResult{
			Found:               false,
			Nick:                nick,
			StartCrawlerSuccess: false,
		})
		return err
	}
	return nil
}

// GetUserInfoFromThunderskill
// TODO: thunderskill启用了cloudflare，暂不爬取
func GetUserInfoFromThunderskill(nick string) error {
	urlTemplate := "https://thunderskill.com/en/stat/%s/export/json"
	url := fmt.Sprintf(urlTemplate, nick)

	c := colly.NewCollector(
		colly.AllowedDomains("thunderskill.com"),
		colly.MaxDepth(1),
		colly.IgnoreRobotsTxt(),
	)

	c.OnResponse(func(r *colly.Response) {
		mp := make(map[string]any)
		_ = json.Unmarshal(r.Body, &mp)
		logging.Info(r.Body)
	})

	c.OnRequest(func(r *colly.Request) {
		logging.Infof("start visiting %s", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		logging.Warnf("visiting %s failed. %s", r.Request.URL.String(), err)
	})

	err := c.Visit(url)
	if err != nil {
		logging.Warn("colly get failed. ", err)
		return err
	}
	return nil
}
