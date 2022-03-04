package service

import (
	"axiangcoding/antonstar/api-system/data"
	"axiangcoding/antonstar/api-system/data/schema"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"time"
)

type VisitItem struct {
	// 客户端生成id
	ClientId string
	// 用户id
	UserId int64
	// 访问页面
	Page      string
	VisitTime time.Time
}

func AddVisit(c *gin.Context, visit VisitItem) error {
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
