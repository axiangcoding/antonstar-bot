package schema

import (
	"axiangcoding/antonstar/api-system/pkg/logging"
	"database/sql"
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

const (
	// UserRoleAdmin 超级管理员
	UserRoleAdmin = "admin"
	// UserRoleManager 维护者
	UserRoleManager = "manager"
	// UserRoleOrdinary 普通用户
	UserRoleOrdinary = "ordinary"
)

const (
	// UserStatusNoVerify 未验证状态
	UserStatusNoVerify = "not_verify"
	// UserStatusNormal 正常状态
	UserStatusNormal = "normal"
	// UserStatusBanned 封禁状态
	UserStatusBanned = "banned"
	// UserStatusMute 禁言状态
	UserStatusMute = "mute"
)

type User struct {
	gorm.Model
	UserId int64 `gorm:"uniqueIndex"`
	// 登录用户名
	UserName string `gorm:"uniqueIndex;size:255"`
	// 用户昵称
	NickName sql.NullString `gorm:"uniqueIndex;size:255"`
	// 头像链接
	AvatarUrl string `gorm:"size:255"`
	// 邮箱
	Email sql.NullString `gorm:"uniqueIndex;size:255"`
	// 电话
	Phone sql.NullString `gorm:"uniqueIndex;size:255"`
	// 邀请码
	InvitedCode string `gorm:"size:255"`
	// 加密后的密码
	Password string `gorm:"size:255"`
	// 用户分配的角色，逗号分割的若干个值
	Roles string `gorm:"size:255"`
	// 用户状态
	Status string `gorm:"size:255"`
	// 达人认证ID
	CertId uint
}

func (u *User) GenerateId() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		logging.Error("generate snowflake id error", err)
		return
	}
	u.UserId = node.Generate().Int64()
}
