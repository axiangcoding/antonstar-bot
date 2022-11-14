package table

import "gorm.io/gorm"

type QQUserConfig struct {
	gorm.Model
	UserId           int64 `gorm:"uniqueIndex;"`
	Banned           *bool
	TodayQueryCount  int
	OneDayQueryLimit int
	TotalQueryCount  int
}

func DefaultUserConfig(userId int64) QQUserConfig {
	// trueVal := true
	falseVal := false
	return QQUserConfig{
		UserId:           userId,
		Banned:           &falseVal,
		TodayQueryCount:  0,
		OneDayQueryLimit: 10,
		TotalQueryCount:  0,
	}
}
