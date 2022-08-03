package service

import (
	"encoding/json"
	"errors"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/service/crawler"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"time"
)

type ScheduleForm struct {
	SendForm         cqhttp.SendGroupMsgForm `json:"send_form"`
	Project          string                  `json:"project,omitempty"`
	Spider           string                  `json:"spider,omitempty"`
	MissionId        string                  `json:"missionId,omitempty"`
	Nick             string                  `json:"nick,omitempty"`
	CallbackEndpoint string                  `json:"callback_endpoint,omitempty"`
}

type ScheduleResult struct {
	NodeName string `json:"node_name,omitempty"`
	Status   string `json:"status,omitempty"`
	JobId    string `json:"jobid,omitempty"`
}

func RequestCrawlerSpider(form ScheduleForm) error {
	url := settings.Config.Service.Crawler.Url + "/schedule.json"
	var result ScheduleResult
	client := resty.New().SetTimeout(time.Second * 20)
	resp, err := client.R().
		SetFormData(map[string]string{
			"project":    form.Project,
			"spider":     form.Spider,
			"mission_id": form.MissionId,
			"nick":       form.Nick,
			// "callback_endpoint": form.CallbackEndpoint,
		}).SetResult(&result).Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return errors.New("request crawler spider error")
	}
	return nil
}

func HandleCrawlerCallback(missionId string, source string, data map[string]any) error {
	var nick string
	if source == crawler.SourceGaijin {
		var gaijinData crawler.GaijinData
		if err := mapstructure.Decode(data, &gaijinData); err != nil {
			return err
		}
		nick = gaijinData.Nick
		// upsert
		_, err := FindGameProfile(nick)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := SaveGameProfile(table.GameUser{Nick: nick}); err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if err := UpdateGameProfile(nick, gaijinData.ToTableGameUser()); err != nil {
			return err
		}
		if err := FinishMission(missionId, table.MissionStatusSuccess, gaijinData); err != nil {
			return err
		}
	} else if source == crawler.SourceThunderSkill {
		var thunderSkillData crawler.ThunderSkillData
		if err := mapstructure.Decode(data, &thunderSkillData); err != nil {
			return err
		}
		nick = thunderSkillData.Nick
		// upsert
		_, err := FindGameProfile(nick)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := SaveGameProfile(table.GameUser{Nick: nick}); err != nil {
					return err
				}
			} else {
				return err
			}
		}
		if err := UpdateGameProfile(nick, thunderSkillData.ToTableGameUser()); err != nil {
			return err
		}
		if err := FinishMission(missionId, table.MissionStatusSuccess, thunderSkillData); err != nil {
			return err
		}
	} else {
		logging.Warnf("no such source %s", source)
	}
	return nil
}

func WaitForCrawlerCallback(missionIds []string) error {
	totalDelay := 60
	duration := 3
	i := 0
	var detailForm ScheduleForm
	for i <= totalDelay {
		time.Sleep(time.Second * time.Duration(duration))
		allDone := true

		for _, id := range missionIds {
			mission, err := FindMission(id)
			if err != nil {
				logging.Warnf("polling find mission failed. %s", err)
				i += duration
				continue
			}
			detail := mission.Detail
			if err := json.Unmarshal([]byte(detail), &detailForm); err != nil {
				logging.Warnf("unmarshal mission.detail error. %s", err)
			}
			allDone = mission.Status == table.MissionStatusSuccess && allDone
		}
		if allDone {
			user, err := FindGameProfile(detailForm.Nick)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					detailForm.SendForm.Message = "未找到该用户，请检查游戏昵称是否正确"
				}
			} else {
				detailForm.SendForm.Message = user.ToDisplayGameUser().ToFriendlyString()
			}
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
