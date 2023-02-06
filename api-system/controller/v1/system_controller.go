package v1

import (
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/gin-gonic/gin"
)

// SystemInfo
// @Summary  System Info
// @Tags     System API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/system/info [get]
func SystemInfo(c *gin.Context) {
	m := map[string]string{
		"version": settings.Config.Version,
	}
	app.Success(c, m)
}
