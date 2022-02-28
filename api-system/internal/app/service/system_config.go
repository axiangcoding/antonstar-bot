package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetRefreshLimit(c *gin.Context) int {
	defaultVal := 100
	config, err := data.TakeSystemConfig(c, schema.ConfigDailyRefreshLimit)
	if err != nil {
		return defaultVal
	}
	val, err := strconv.Atoi(config.Value)
	if err != nil {
		return defaultVal
	}
	return val
}
