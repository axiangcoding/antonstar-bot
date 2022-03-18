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
	// TS街机效率值
	TsABRate float64
	// TS历史效率值
	TsRBRate float64
	// TS全真效率值
	TsSBRate float64
	// 安东星街机效率值
	AsABRate float64
	// 安东星历史效率值
	AsRBRate float64
	// 安东星全真效率值
	AsSBRate float64
	// 是否被封禁
	Banned bool
}
