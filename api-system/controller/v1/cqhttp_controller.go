package v1

import (
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/axiangcoding/antonstar-bot/entity/e"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/service"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/axiangcoding/antonstar-bot/tool"
	"github.com/gin-gonic/gin"
)

// CqHttpReceiveEvent
// @Summary   接收cqhttp的事件
// @Tags      CQHttp API
// @Param     event  body      object       true  "cqhttp event"
// @Success   200    {object}  app.ApiJson  ""
// @Router    /v1/cqhttp/receive/event [post]
// @Security  AppToken
func CqHttpReceiveEvent(c *gin.Context) {
	var m map[string]any
	err := c.ShouldBindJSON(&m)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	cp := c.Copy()
	tool.GoWithRecover(func() {
		if err := service.HandleCqHttpEvent(cp, m); err != nil {
			logging.Errorf("async handle cqhttp event failed. %s", err)
		}
	})

	app.Success(c, nil)
}

// CqHttpStatus
// @Summary  获取cqhttp的最新状态
// @Tags     CQHttp API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/cqhttp/status [get]
func CqHttpStatus(c *gin.Context) {
	defaultSelfId := settings.Config.Service.CqHttp.SelfQQ
	status, err := service.GetCqHttpStatus(c, defaultSelfId)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, status)
}
