package schema

import "gorm.io/gorm"

type SystemConfig struct {
	gorm.Model
	// 键名
	Key string `gorm:"uniqueIndex;size:255"`
	// 值名
	Value string `gorm:"size:255"`
}

const (
	ConfigDailyRefreshLimit = "daily_refresh_limit"
)
