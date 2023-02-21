package v1

import (
	"github.com/axiangcoding/antonstar-bot/internal/entity/app"
	"github.com/axiangcoding/antonstar-bot/internal/entity/e"
	"github.com/axiangcoding/antonstar-bot/internal/service"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/axiangcoding/antonstar-bot/setting"
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
)

// CqHttpReceiveEvent
// @Summary   接收cqhttp的事件
// @Tags      CQHttp API
// @Param     event  body      object       true  "cqhttp event"
// @Success   200    {object}  app.ApiJson  ""
// @Router    /v1/cqhttp/receive/event [post]
// @Security  CqhttpSelfID
// @Security  CqhttpSignature
func CqHttpReceiveEvent(c *gin.Context) {
	var m map[string]any
	err := c.ShouldBindJSON(&m)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	cp := c.Copy()
	if err := ants.Submit(func() {
		if err := service.HandleCqHttpEvent(cp, m); err != nil {
			logging.L().Error("async handle cqhttp event failed. %s", logging.Error(err))
		}
	}); err != nil {
		logging.L().Error("ant submit error.", logging.Error(err))
	}
	app.Success(c, nil)
}

// CqHttpStatus
// @Summary  获取cqhttp的最新状态
// @Tags     CQHttp API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/cqhttp/status [get]
func CqHttpStatus(c *gin.Context) {
	defaultSelfId := setting.C().App.Service.CqHttp.SelfQQ
	status, err := service.GetCqHttpStatus(c, defaultSelfId)
	logging.L().Info("", logging.Any("status", status))
	mp := map[string]any{
		"app_enabled": status.Status.AppEnabled,
		"app_good":    status.Status.AppGood,
		"online":      status.Status.Online,
		"plugin_good": status.Status.PluginsGood,
	}
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, mp)
}
