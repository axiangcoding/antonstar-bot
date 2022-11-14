package cron

import (
	"fmt"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service"
	"github.com/axiangcoding/ax-web/service/bot"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/robfig/cron/v3"
)

func Setup() {
	c := cron.New()
	addJob(c)
	c.Start()
}

func addJob(c *cron.Cron) {
	if _, err := c.AddFunc("@every 5m", CheckRoomLiving); err != nil {
		logging.Fatalf("Add cron job CheckRoomLiving failed. %s", err)
	}
	if _, err := c.AddFunc("@daily", RefreshUserTodayCount); err != nil {
		logging.Fatalf("Add cron job RefreshUserTodayCount failed. %s", err)
	}
	logging.Info("All cron job add success")
}

func CheckRoomLiving() {
	enableCheckBiliRoom := true
	qcs, err := service.GetGroupConfigWithCondition(table.QQGroupConfig{EnableCheckBiliRoom: &enableCheckBiliRoom})
	if err != nil {
		logging.Warn(err)
		return
	}
	for _, qc := range qcs {
		var sgmf cqhttp.SendGroupMsgForm
		sgmf.GroupId = qc.GroupId
		info, err := service.GetBiliBiliRoomInfo(qc.BindBiliRoomId)
		if err != nil {
			logging.Warn(err)
			continue
		}
		if info.Data.LiveStatus == 1 {
			exist := service.ExistBiliRoomFlag(qc.GroupId, qc.BindBiliRoomId)
			if !exist {
				url := fmt.Sprintf("https://live.bilibili.com/%d", qc.BindBiliRoomId)
				config, _ := service.FindGroupConfig(qc.GroupId)
				sgmf.Message = fmt.Sprintf(bot.SelectStaticMessage(config.MessageTemplate).CommonResp.LiveBroadcast, info.Data.Title, url)
				service.MustSendGroupMsg(sgmf)
			}
			service.MustPutBiliRoomFlag(qc.GroupId, qc.BindBiliRoomId)
		}
	}

}

func RefreshUserTodayCount() {
	if err := service.ResetAllUserConfigTodayCount(); err != nil {
		logging.Error("reset all qq user today_query_count failed. ", err)
	} else {
		logging.Info("reset all qq user today_query_count to 0")
	}

	if err := service.ResetAllGroupConfigTodayCount(); err != nil {
		logging.Error("reset all qq group today_query_count failed. ", err)
	} else {
		logging.Info("reset all qq group today_query_count to 0")
	}
}
