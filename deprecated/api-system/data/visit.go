package data

import (
	"axiangcoding/antonstar/api-system/data/schema"
	"axiangcoding/antonstar/api-system/logging"
	"context"
	"time"
)

func SaveVisit(c context.Context, visit schema.Visit) error {
	err := GetDB().Save(&visit).Error
	return err
}

func CountVisit(c context.Context, timestamp time.Time) int64 {
	var count int64
	model := GetDB().Model(&schema.Visit{})
	if !timestamp.IsZero() {
		model.Where("to_days(created_at) = to_days(?)", timestamp)
	}
	err := model.Count(&count).Error
	if err != nil {
		logging.Error(err)
	}
	return count
}
