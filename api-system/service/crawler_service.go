package service

import (
	"errors"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/service/crawler"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/go-resty/resty/v2"
	"github.com/mitchellh/mapstructure"
	"time"
)

type ScheduleForm struct {
	Project          string `json:"project,omitempty"`
	Spider           string `json:"spider,omitempty"`
	MissionId        string `json:"missionId,omitempty"`
	Nick             string `json:"nick,omitempty"`
	CallbackEndpoint string `json:"callback_endpoint,omitempty"`
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

func HandleCrawlerCallback(missionId string, data map[string]any) error {
	var gaijinData crawler.GaijinData
	if err := mapstructure.Decode(data, &gaijinData); err != nil {
		return err
	}
	if err := FinishMission(missionId, table.MissionStatusSuccess, gaijinData); err != nil {
		return err
	}
	return nil
}
