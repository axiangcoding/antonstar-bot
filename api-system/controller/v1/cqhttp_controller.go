package v1

import (
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/entity/e"
	"github.com/axiangcoding/ax-web/service"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/gin-gonic/gin"
)

type ReceiveForm struct {
}

// CqHttpReceiveEvent
// @Summary  receive event from cqhttp service
// @Tags     CQHttp API
// @Param    accessToken  query     string       true  "accessToken"
// @Param    param        body      object       true  "getParam"
// @Success  200          {object}  app.ApiJson  ""
// @Router   /v1/cqhttp/receive/event [post]
func CqHttpReceiveEvent(c *gin.Context) {
	accessToken := c.Query("accessToken")
	if accessToken != "defaultAccessToken" {
		app.Unauthorized(c, e.TokenNotValid)
		return
	}
	var m map[string]any
	err := c.ShouldBindJSON(&m)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	if err := service.HandleCqHttpEvent(c, m); err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, nil)
}

// CqHttpStatus
// @Summary  get cqhttp service status
// @Tags     CQHttp API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/cqhttp/status [get]
func CqHttpStatus(c *gin.Context) {
	defaultSelfId := settings.Config.CqHttp.SelfQQ
	status, err := service.GetCqHttpStatus(c, defaultSelfId)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, status)
}
