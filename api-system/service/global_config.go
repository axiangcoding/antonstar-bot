package service

import (
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
)

func FindGlobalConfig(key string) (*table.GlobalConfig, error) {
	db := data.GetDB()
	var find table.GlobalConfig
	if err := db.Where(table.GlobalConfig{Key: key}).Take(&find).Error; err != nil {
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
