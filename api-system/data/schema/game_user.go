package schema

import (
	"gorm.io/gorm"
	"time"
)

type GameUser struct {
	gorm.Model
	// 游戏昵称
	Nick string `gorm:"uniqueIndex;size:255"`
	// 联队
	Clan string `gorm:"size:255"`
	// 联队地址
	ClanUrl string `gorm:"size:255"`
	// 注册日期
	RegisterDate time.Time
	// 游戏等级
	Level int
	// 称号
	Title string `gorm:"size:255"`
	// 是否被封禁
	Banned bool
}
