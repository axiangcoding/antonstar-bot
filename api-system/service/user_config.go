package service

import (
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
)

func FindUserConfig(userId int64) (*table.QQUserConfig, error) {
	db := data.GetDB()
	var find table.QQUserConfig
	if err := db.Where(table.QQUserConfig{UserId: userId}).Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func MustFindUserConfig(userId int64) *table.QQUserConfig {
	config, err := FindUserConfig(userId)
	if err != nil {
		logging.Warn(err)
	}
	return config
}

func SaveUserConfig(gc table.QQUserConfig) error {
	db := data.GetDB()
	if err := db.Save(&gc).Error; err != nil {
		return err
	}
	return nil
}

func CheckUserTodayQueryLimit(userId int64) (bool, int, int) {
	config := MustFindUserConfig(userId)
	if config == nil {
		return true, 0, 0
	}
	return config.TodayQueryCount >= config.OneDayQueryLimit, config.TodayQueryCount, config.OneDayQueryLimit
}

func MustAddUserConfigTodayQueryCount(userId int64, count int) {
	config := MustFindUserConfig(userId)
	config.TodayQueryCount += count
	err := SaveUserConfig(*config)
	if err != nil {
		logging.Warn(err)
	}
}

func MustAddUserConfigTotalQueryCount(userId int64, count int) {
	config := MustFindUserConfig(userId)
	config.TotalQueryCount += count
	err := SaveUserConfig(*config)
	if err != nil {
		logging.Warn(err)
	}
}

func ResetAllUserConfigTodayQueryCount() error {
	db := data.GetDB()
	if err := db.Model(&table.QQUserConfig{}).
		Select("today_query_count").Where("1=1").
		Updates(table.QQUserConfig{TodayQueryCount: 0}).Error; err != nil {
		return err
	}
	return nil
}
