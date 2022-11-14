package table

import "gorm.io/gorm"

type QQUserConfig struct {
	gorm.Model
	UserId           int64 `gorm:"uniqueIndex;"`
	Banned           *bool
	TodayQueryCount  int
	OneDayQueryLimit int
	TotalQueryCount  int
	TodayUsageCount  int
	OneDayUsageLimit int
	TotalUsageCount  int
}

func DefaultUserConfig(userId int64) QQUserConfig {
	// trueVal := true
	falseVal := false
	return QQUserConfig{
		UserId:           userId,
		Banned:           &falseVal,
		TodayQueryCount:  0,
		OneDayQueryLimit: 5,
		TotalQueryCount:  0,
		TodayUsageCount:  0,
		OneDayUsageLimit: 50,
		TotalUsageCount:  0,
	}
}
