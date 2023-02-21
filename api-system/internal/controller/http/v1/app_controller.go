package v1

import (
	"github.com/axiangcoding/antonstar-bot/internal/entity/app"
	"github.com/axiangcoding/antonstar-bot/setting"
	"github.com/gin-gonic/gin"
)

// AppInfo
// @Summary  获取应用信息
// @Tags     App API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/app/info [get]
func AppInfo(c *gin.Context) {
	m := map[string]string{
		"version": setting.C().App.Version,
	}
	app.Success(c, m)
}
