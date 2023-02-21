package table

import (
	"fmt"
	"github.com/axiangcoding/antonstar-bot/internal/data/display"
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
	Banned *bool
	// 注册日期
	RegisterDate time.Time
	// 称号
	Title string `gorm:"size:255"`
	// 游戏等级
	Level  int
	StatAb UserStat `gorm:"embedded;embeddedPrefix:stat_ab_"`
	StatRb UserStat `gorm:"embedded;embeddedPrefix:stat_rb_"`
	StatSb UserStat `gorm:"embedded;embeddedPrefix:stat_sb_"`

	GroundRateAb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_ab_"`
	GroundRateRb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_rb_"`
	GroundRateSb GroundRate `gorm:"embedded;embeddedPrefix:rate_ground_sb_"`

	AviationRateAb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_ab_"`
	AviationRateRb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_rb_"`
	AviationRateSb AviationRate `gorm:"embedded;embeddedPrefix:rate_aviation_sb_"`

	FleetRateAb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_ab_"`
	FleetRateRb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_rb_"`
	FleetRateSb FleetRate `gorm:"embedded;embeddedPrefix:rate_fleet_sb_"`

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

type UserStat struct {
	TotalMission         int
	WinRate              float64
	GroundDestroyCount   int
	FleetDestroyCount    int
	GameTime             string
	AviationDestroyCount int
	WinCount             int
	SliverEagleEarned    int64
	DeadCount            int
}

type GroundRate struct {
	GameCount              int
	GroundVehicleGameCount int
	TDGameCount            int
	HTGameCount            int
	SPAAGameCount          int
	GameTime               string
	GroundVehicleGameTime  string
	TDGameTime             string
	HTGameTime             string
	SPAAGameTime           string
	TotalDestroyCount      int
	AviationDestroyCount   int
	GroundDestroyCount     int
	FleetDestroyCount      int
}

type AviationRate struct {
	GameCount            int
	FighterGameCount     int
	BomberGameCount      int
	AttackerGameCount    int
	GameTime             string
	FighterGameTime      string
	BomberGameTime       string
	AttackerGameTime     string
	TotalDestroyCount    int
	AviationDestroyCount int
	GroundDestroyCount   int
	FleetDestroyCount    int
}

type FleetRate struct {
	GameCount               int
	FleetGameCount          int
	TorpedoBoatGameCount    int
	GunboatGameCount        int
	TorpedoGunboatGameCount int
	SubmarineHuntGameCount  int
	DestroyerGameCount      int
	NavyBargeGameCount      int
	GameTime                string
	FleetGameTime           string
	TorpedoBoatGameTime     string
	GunboatGameTime         string
	TorpedoGunboatGameTime  string
	SubmarineHuntGameTime   string
	DestroyerGameTime       string
	NavyBargeGameTime       string
	TotalDestroyCount       int
	AviationDestroyCount    int
	GroundDestroyCount      int
	FleetDestroyCount       int
}

func (u GameUser) ToDisplayGameUser() display.GameUser {
	zone := time.FixedZone("CST", 8*3600)
	return display.GameUser{
		CreatedAt:      u.CreatedAt.In(zone).Format("2006-01-02 15:04:05"),
		UpdatedAt:      u.UpdatedAt.In(zone).Format("2006-01-02 15:04:05"),
		Nick:           u.Nick,
		Clan:           u.Clan,
		ClanUrl:        u.ClanUrl,
		RegisterDate:   u.RegisterDate.Format("2006-01-02"),
		Level:          u.Level,
		Title:          u.Title,
		TsABRate:       u.TsABRate,
		TsRBRate:       u.TsRBRate,
		TsSBRate:       u.TsSBRate,
		AsABRate:       u.AsABRate,
		AsRBRate:       u.AsRBRate,
		AsSBRate:       u.AsSBRate,
		Banned:         *u.Banned,
		StatSb:         convertToStat(u.StatSb),
		StatAb:         convertToStat(u.StatAb),
		StatRb:         convertToStat(u.StatRb),
		GroundRateAb:   convertToGroundRate(u.GroundRateAb),
		GroundRateRb:   convertToGroundRate(u.GroundRateRb),
		GroundRateSb:   convertToGroundRate(u.GroundRateSb),
		AviationRateAb: convertToAviationRate(u.AviationRateAb),
		AviationRateRb: convertToAviationRate(u.AviationRateRb),
		AviationRateSb: convertToAviationRate(u.AviationRateSb),
		FleetRateAb:    convertToFleetRate(u.FleetRateAb),
		FleetRateRb:    convertToFleetRate(u.FleetRateRb),
		FleetRateSb:    convertToFleetRate(u.FleetRateSb),
	}
}

func convertToStat(stat UserStat) display.UserStat {
	var kd float64
	if stat.DeadCount != 0 {
		kd = float64(stat.GroundDestroyCount+stat.FleetDestroyCount+stat.AviationDestroyCount) / float64(stat.DeadCount)
	}
	return display.UserStat{
		TotalMission:         stat.TotalMission,
		WinRate:              fmt.Sprintf("%.0f%%", stat.WinRate*100),
		GroundDestroyCount:   stat.GroundDestroyCount,
		FleetDestroyCount:    stat.FleetDestroyCount,
		GameTime:             stat.GameTime,
		AviationDestroyCount: stat.AviationDestroyCount,
		WinCount:             stat.WinCount,
		SliverEagleEarned:    stat.SliverEagleEarned,
		DeadCount:            stat.DeadCount,
		Kd:                   fmt.Sprintf("%.2f", kd),
	}
}

func convertToAviationRate(rate AviationRate) display.UserRate {
	ka := float64(rate.TotalDestroyCount) / float64(rate.GameCount)
	return display.UserRate{
		Ka: fmt.Sprintf("%.2f", ka),
	}
}

func convertToGroundRate(rate GroundRate) display.UserRate {
	ka := float64(rate.TotalDestroyCount) / float64(rate.GameCount)
	return display.UserRate{
		Ka: fmt.Sprintf("%.2f", ka),
	}
}

func convertToFleetRate(rate FleetRate) display.UserRate {
	ka := float64(rate.TotalDestroyCount) / float64(rate.GameCount)
	return display.UserRate{
		Ka: fmt.Sprintf("%.2f", ka),
	}
}
