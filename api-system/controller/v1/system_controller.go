package v1

import (
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/gin-gonic/gin"
)

// SystemInfo
// @Summary  获取系统信息
// @Tags     System API
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/system/info [get]
func SystemInfo(c *gin.Context) {
	m := map[string]string{}
	app.Success(c, m)
}
