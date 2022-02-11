package schema

import "gorm.io/gorm"

// UserCert 用户认证
type UserCert struct {
	gorm.Model
	// 认证信息
	Info string `gorm:"size:255"`
	// 认证类型
	Type string `gorm:"size:255"`
	// 是否吊销认证
	Revoked bool
	// 授予人
	GrantBy string `gorm:"size:255"`
}
