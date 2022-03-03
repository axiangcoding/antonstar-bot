package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/pkg/app"
	"github.com/gin-gonic/gin"
)

// SystemInfo
// @Summary  System Info
// @Tags     System
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/system/info [get]
func SystemInfo(c *gin.Context) {
	m := map[string]string{
		"version": conf.Config.Version,
	}
	app.Success(c, m)
}
