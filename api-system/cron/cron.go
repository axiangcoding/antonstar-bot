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
	_, err := c.AddFunc("@every 2m", CheckRoomLiving)
	if err != nil {
		logging.Fatalf("Add cron job failed. %s", err)
	}
	logging.Info("Add cron job success")
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
				sgmf.Message = fmt.Sprintf(bot.RespRoomIsLiving, info.Data.Title, url)
				service.MustSendGroupMsg(sgmf)
			}
			service.MustPutBiliRoomFlag(qc.GroupId, qc.BindBiliRoomId)
		}
	}

}
