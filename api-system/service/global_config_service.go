package service

import (
	"github.com/axiangcoding/ax-web/data/dal"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
)

func FindGlobalConfig(key string) (*table.GlobalConfig, error) {
	take, err := dal.Q.GlobalConfig.Where(dal.GlobalConfig.Key.Eq(key)).Take()
	if err != nil {
		return nil, err
	}
	return take, nil
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
	if err := dal.Q.GlobalConfig.Save(config); err != nil {
		logging.Warn(err)
	}
}
