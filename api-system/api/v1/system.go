package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/cache"
	"github.com/gin-gonic/gin"
)

// SystemInfo
// @Summary   System Info
// @Tags      System
// @Success   200  {object}  app.ApiJson  ""
// @Router    /v1/system/info [get]
// @Security  ApiKeyAuth
func SystemInfo(c *gin.Context) {
	db := data.GetDB()
	s, _ := db.DB()
	stats := cache.GetRedis().PoolStats()
	app.Success(c, map[string]interface{}{
		"mysql-stats": s.Stats(),
		"redis-stats": stats})
}
