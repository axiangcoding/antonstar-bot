package service

import (
	"encoding/json"
	"errors"
	"github.com/axiangcoding/antonstar-bot/internal/data/table"
	"github.com/axiangcoding/antonstar-bot/pkg/cqhttp"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
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
			logging.L().Warn("polling find mission failed", logging.Error(err))
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
	cqhttp.MustSendGroupMsg(detailForm.SendForm)
	return nil
}
