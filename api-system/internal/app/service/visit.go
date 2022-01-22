package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"time"
)

func AddVisit(c *gin.Context, visit entity.AddVisit) error {
	request := c.Request
	ua := user_agent.New(request.Header.Get("User-Agent"))
	name, version := ua.Browser()
	v := schema.Visit{
		UserId:         visit.UserId,
		ClientId:       visit.ClientId,
		ClientIp:       c.ClientIP(),
		VisitPath:      visit.Page,
		BrowserName:    name,
		BrowserVersion: version,
		Bot:            ua.Bot(),
		Platform:       ua.Platform(),
		OS:             ua.OS(),
		VisitTime:      visit.VisitTime,
	}
	err := data.SaveVisit(c, v)
	if err != nil {
		return err
	}
	return nil
}

func CountVisit(c *gin.Context, timestamp time.Time) map[string]interface{} {
	visit := data.CountVisit(c, timestamp)
	return map[string]interface{}{
		"count": visit,
	}
}
