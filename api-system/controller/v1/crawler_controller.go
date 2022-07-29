package v1

import (
	"encoding/json"
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/entity/e"
	"github.com/axiangcoding/ax-web/service"
	"github.com/gin-gonic/gin"
)

type CrawlerReceiveForm struct {
	MissionId   string `json:"mission_id,omitempty" binding:"required"`
	CrawlerData string `json:"crawler_data,omitempty" binding:"required"`
	Source      string `json:"source" binding:"required"`
}

// ReceiveCrawlerCallback
// @Summary   接收爬虫的回调
// @Tags      Crawler API
// @Param     form  body      CrawlerReceiveForm  true  "form"
// @Success   200   {object}  app.ApiJson         ""
// @Router    /v1/crawler/callback [post]
// @Security  AppToken
func ReceiveCrawlerCallback(c *gin.Context) {
	var form CrawlerReceiveForm
	if err := c.ShouldBindJSON(&form); err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}

	var data map[string]any
	if err := json.Unmarshal([]byte(form.CrawlerData), &data); err != nil {
		app.BizFailed(c, e.RequestParamsNotValid, err)
		return
	}

	if err := service.HandleCrawlerCallback(form.MissionId, data); err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, form)
}
