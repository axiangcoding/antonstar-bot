package data

import (
	"axiangcoding/antonstar/api-system/data/schema"
	"context"
)

func SaveSystemConfig(c context.Context, systemConfig schema.SystemConfig) error {
	err := GetDB().Save(&systemConfig).Error
	return err
}

func TakeSystemConfig(c context.Context, key string) (schema.SystemConfig, error) {
	var findItem schema.SystemConfig
	find := GetDB().Where(&schema.SystemConfig{
		Key: key,
	}).Take(&findItem)
	return findItem, find.Error
}
