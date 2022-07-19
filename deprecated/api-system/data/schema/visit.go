package schema

import (
	"gorm.io/gorm"
	"time"
)

type Visit struct {
	gorm.Model
	// 如果是登录用户，带user_id
	UserId int64
	// 如果没登录，传client_id
	ClientId string `gorm:"size:255"`
	// 访问的ip地址
	ClientIp string `gorm:"size:255"`
	// 访问的页面
	VisitPath string `gorm:"size:255"`
	// 是否是机器人
	Bot bool
	// 浏览器名称
	BrowserName string `gorm:"size:255"`
	// 浏览器版本
	BrowserVersion string `gorm:"size:255"`
	// 架构平台
	Platform string `gorm:"size:255"`
	// 操作系统
	OS string `gorm:"size:255"`
	// 访问时间
	VisitTime time.Time
}
