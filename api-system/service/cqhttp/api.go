package cqhttp

import (
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/go-resty/resty/v2"
	"time"
)

// MustSendGroupMsg
// https://docs.go-cqhttp.org/api/#%E5%A4%84%E7%90%86%E5%8A%A0%E7%BE%A4%E8%AF%B7%E6%B1%82-%E9%82%80%E8%AF%B7
func MustSendGroupMsg(form SendGroupMsgForm) {
	url := settings.C().App.Service.CqHttp.Url + "/send_group_msg"
	client := resty.New().SetTimeout(time.Second * 20)
	var commonResp CommonResponse
	resp, err := client.R().SetHeader("Content-cardType", "application/json").
		SetBody(map[string]any{
			"message":  form.MessagePrefix + form.Message,
			"group_id": form.GroupId,
		}).SetResult(&commonResp).Post(url)
	if err != nil {
		logging.L().Error("send group message error", logging.Error(err))
	}
	if resp.IsError() {
		logging.L().Warn("post error",
			logging.Any("url", url),
			logging.Any("statusCode", resp.StatusCode()),
			logging.Any("resp", resp.String()))
	}
	if commonResp.Status == "failed" {
		logging.L().Error("send message failed", logging.Any("resp", commonResp))
	}
}

// MustAcceptInviteToGroup
// https://docs.go-cqhttp.org/api/#%E5%A4%84%E7%90%86%E5%8A%A0%E7%BE%A4%E8%AF%B7%E6%B1%82-%E9%82%80%E8%AF%B7
func MustAcceptInviteToGroup(flag string, subType string, approve bool, reason string) {
	url := settings.C().App.Service.CqHttp.Url + "/set_group_add_request"
	client := resty.New().SetTimeout(time.Second * 20)
	resp, err := client.R().SetHeader("Content-cardType", "application/json").
		SetBody(map[string]any{
			"flag":     flag,
			"sub_type": subType,
			"approve":  approve,
			"reason":   reason,
		}).Post(url)
	if err != nil {
		logging.L().Error("send group add request error", logging.Error(err))
	}
	if resp.IsError() {
		logging.L().Warn("post error",
			logging.Any("url", url),
			logging.Any("statusCode", resp.StatusCode()),
			logging.Any("resp", resp.String()))
	}
}
