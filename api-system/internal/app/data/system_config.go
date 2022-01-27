package data

import (
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"context"
	"errors"
	"gorm.io/gorm"
)

func SaveSystemConfig(c context.Context, systemConfig schema.SystemConfig) error {
	err := GetDB().Save(&systemConfig).Error
	return err
}

func FindSystemConfig(c context.Context, key string) (schema.SystemConfig, error) {
	var findItem schema.SystemConfig
	find := GetDB().Where(&schema.SystemConfig{
		Key: key,
	}).Take(&findItem)
	if errors.Is(find.Error, gorm.ErrRecordNotFound) {
		return findItem, find.Error
	}
	return findItem, nil
}
