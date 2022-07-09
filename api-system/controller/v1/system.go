package v1

import (
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/gin-gonic/gin"
	"time"
)

// SystemInfo
// @Summary  System Info
// @Tags     System API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/system/info [get]
func SystemInfo(c *gin.Context) {
	time.Sleep(time.Second * 6)
	m := map[string]string{
		"version": settings.Config.Version,
	}
	app.Success(c, m)
}
