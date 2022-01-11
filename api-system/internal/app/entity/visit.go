package entity

import "time"

type AddVisit struct {
	// 客户端生成id
	ClientId string
	// 用户id
	UserId int64
	// 访问页面
	Page      string
	VisitTime time.Time
}
