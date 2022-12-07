package service

import (
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
)

func FindGlobalConfig(key string) (*table.GlobalConfig, error) {
	db := data.GetDB()
	var find table.GlobalConfig
	if err := db.Where(table.GlobalConfig{Key: key}, "key").Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func MustFindGlobalConfig(key string) *table.GlobalConfig {
	config, err := FindGlobalConfig(key)
	if err != nil {
		logging.Warn(err)
	}
	return config
}

func IsStopGlobalQuery() bool {
	config := MustFindGlobalConfig(table.ConfigStopQuery)
	if config == nil {
		return true
	} else {
		return config.Value == "true"
	}
}

func IsStopAllResponse() bool {
	config := MustFindGlobalConfig(table.ConfigStopAllResponse)
	if config == nil {
		return false
	} else {
		return config.Value == "true"
	}
}

func MustUpsertGlobalConfig(key string, value string) {
	config := MustFindGlobalConfig(key)
	if config == nil {
		config = &table.GlobalConfig{Key: key, Value: value}
	} else {
		config.Value = value
	}
	db := data.GetDB()
	db.Save(&config)
}
