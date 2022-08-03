package table

import (
	"github.com/axiangcoding/ax-web/data/display"
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
	// 是否被封禁
	Banned bool
	// 注册日期
	RegisterDate time.Time
	// 称号
	Title string `gorm:"size:255"`
	// 游戏等级
	Level int
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
}

func (u GameUser) ToDisplayGameUser() display.GameUser {
	var bannedStr string
	if u.Banned {
		bannedStr = "是"
	} else {
		bannedStr = "否"
	}
	zone := time.FixedZone("CST", 8*3600)
	return display.GameUser{
		CreatedAt:    u.CreatedAt.In(zone).Format("2006-01-02 15:04:05"),
		UpdatedAt:    u.UpdatedAt.In(zone).Format("2006-01-02 15:04:05"),
		Nick:         u.Nick,
		Clan:         u.Clan,
		ClanUrl:      u.ClanUrl,
		RegisterDate: u.RegisterDate.Format("2006-01-02"),
		Level:        u.Level,
		Title:        u.Title,
		TsABRate:     u.TsABRate,
		TsRBRate:     u.TsRBRate,
		TsSBRate:     u.TsSBRate,
		AsABRate:     u.AsABRate,
		AsRBRate:     u.AsRBRate,
		AsSBRate:     u.AsSBRate,
		Banned:       bannedStr,
	}
}
