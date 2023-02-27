package cron

import (
	"fmt"
	"github.com/axiangcoding/antonstar-bot/internal/data/table"
	"github.com/axiangcoding/antonstar-bot/internal/service"
	"github.com/axiangcoding/antonstar-bot/pkg/bilibili"
	"github.com/axiangcoding/antonstar-bot/pkg/bot"
	"github.com/axiangcoding/antonstar-bot/pkg/cqhttp"
	"github.com/axiangcoding/antonstar-bot/pkg/crawler"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/robfig/cron/v3"
)

func InitCronJob() {
	c := cron.New()
	addJob(c)
	c.Start()
}

func addJob(c *cron.Cron) {
	if _, err := c.AddFunc("@every 5m", CheckRoomLiving); err != nil {
		logging.L().Fatal("add cron job CheckRoomLiving failed", logging.Error(err))
	}
	if _, err := c.AddFunc("@daily", RefreshUserTodayCount); err != nil {
		logging.L().Fatal("add cron job RefreshUserTodayCount failed", logging.Error(err))
	}
	if _, err := c.AddFunc("@every 2m", func() {
		CheckWTNewsUpdate("en")
		CheckWTNewsUpdate("zh")
	}); err != nil {
		logging.L().Fatal("add cron job CheckWTNewsUpdate failed", logging.Error(err))
	}
	logging.L().Info("all cron job add success")
}

func CheckRoomLiving() {
	if service.IsStopAllResponse() {
		return
	}
	qcs, err := service.GetEnableCheckBiliRoomGroupConfig(true)
	if err != nil {
		logging.L().Warn("get group config checkbilibiliroom failed", logging.Error(err))
		return
	}
	for _, qc := range qcs {
		var sgmf cqhttp.SendGroupMsgForm
		sgmf.GroupId = qc.GroupId
		info, err := bilibili.GetBiliBiliRoomInfo(qc.BindBiliRoomId)
		if err != nil {
			logging.L().Error("get bilibiliroom info failed", logging.Error(err))
			continue
		}
		if info.Data.LiveStatus == 1 {
			exist := service.ExistBiliRoomFlag(qc.GroupId, qc.BindBiliRoomId)
			if !exist {
				url := fmt.Sprintf("https://live.bilibili.com/%d", qc.BindBiliRoomId)
				config, _ := service.FindGroupConfig(qc.GroupId)
				sgmf.Message = fmt.Sprintf(bot.SelectStaticMessage(config.MessageTemplate).CommonResp.LiveBroadcast, info.Data.Title, url)
				cqhttp.MustSendGroupMsg(sgmf)
			}
			service.MustPutBiliRoomFlag(qc.GroupId, qc.BindBiliRoomId)
		}
	}

}

func RefreshUserTodayCount() {
	if err := service.ResetAllUserConfigTodayCount(); err != nil {
		logging.L().Error("reset all qq user today_query_count failed. ", logging.Error(err))
	} else {
		logging.L().Info("reset all qq user today_query_count to 0")
	}

	if err := service.ResetAllGroupConfigTodayCount(); err != nil {
		logging.L().Error("reset all qq group today_query_count failed. ", logging.Error(err))
	} else {
		logging.L().Info("reset all qq group today_query_count to 0")
	}
}

func CheckWTNewsUpdate(region string) {
	if err := crawler.GetFirstPageNewsFromWTOfficial(region, func(news []table.GameNew) {
		for _, item := range news {
			found := service.MustFindGameNewByLink(item.Link)
			if found == nil {
				service.MustSaveGameNew(&item)
				if service.IsStopAllResponse() {
					return
				}
				// 向配置了的群发送消息
				qcs, err := service.GetEnableCheckWTNew(true)
				if err != nil {
					logging.L().Warn("get group config enable_check_wt_new failed", logging.Error(err))
					return
				}
				for _, qc := range qcs {
					var sgmf cqhttp.SendGroupMsgForm
					sgmf.GroupId = qc.GroupId
					sgmf.MessagePrefix = ""
					sgmf.Message = item.ToDisplayGameUser().ToFriendlyString()
					cqhttp.MustSendGroupMsg(sgmf)
				}
			}
		}
	}); err != nil {
		logging.L().Error("check wt news failed. ", logging.Error(err))
	}
}
