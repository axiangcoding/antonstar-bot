package cqhttp

import (
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/go-resty/resty/v2"
	"time"
)

type SendGroupMsgForm struct {
	GroupId int64  `json:"group_id,omitempty"`
	Message string `json:"message,omitempty"`
}

type CommonResponse struct {
	Status  string `json:"status,omitempty"`
	Retcode int    `json:"retcode,omitempty"`
}

func SendGroupMsg(form SendGroupMsgForm) (*CommonResponse, error) {
	// TODO: 在go 1.19中使用新特性进行附加
	url := settings.Config.Service.CqHttp.Url + "/send_group_msg"
	client := resty.New().SetTimeout(time.Second * 20)
	var commonResp CommonResponse
	resp, err := client.R().SetHeader("Content-Type", "application/json").
		SetBody(form).SetResult(&commonResp).Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		logging.Warnf("post %s error. code=%d, message=%s", url, resp.StatusCode(), resp.String())
		return nil, nil
	}
	return &commonResp, nil
}
