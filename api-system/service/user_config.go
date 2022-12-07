package service

import (
	"errors"
	"github.com/axiangcoding/ax-web/cache"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/go-redis/redis/v8"
	"time"
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

func UpdateUserConfigBindingGameNick(userId int64, gameNick *string) error {
	config := MustFindUserConfig(userId)
	if config == nil {
		return errors.New("qq_user_config not exist")
	} else {
		config.BindingGameNick = gameNick
	}
	db := data.GetDB()
	db.Save(&config)
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

func CheckUserTodayUsageLimit(userId int64) (bool, int, int) {
	config := MustFindUserConfig(userId)
	if config == nil {
		return true, 0, 0
	}
	return config.TodayUsageCount >= config.OneDayUsageLimit, config.TodayUsageCount, config.OneDayUsageLimit
}

func MustAddUserConfigTodayUsageCount(userId int64, count int) {
	config := MustFindUserConfig(userId)
	config.TodayUsageCount += count
	err := SaveUserConfig(*config)
	if err != nil {
		logging.Warn(err)
	}
}

func MustAddUserConfigTotalUsageCount(userId int64, count int) {
	config := MustFindUserConfig(userId)
	config.TotalUsageCount += count
	err := SaveUserConfig(*config)
	if err != nil {
		logging.Warn(err)
	}
}

func ResetAllUserConfigTodayCount() error {
	db := data.GetDB()
	if err := db.Model(&table.QQUserConfig{}).
		Select("today_query_count", "today_usage_count").Where("1=1").
		Updates(table.QQUserConfig{TodayQueryCount: 0, TodayUsageCount: 0}).Error; err != nil {
		return err
	}
	return nil
}

func ExistUserUsageLimitFlag(userId int64) bool {
	client := cache.GetClient()
	key := cache.GenerateUserUsageLimitCacheKey(userId)
	if _, err := client.Get(c, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return false
		}
		logging.Warn(err)
		return false
	}
	return true
}

func MustPutUserUsageLimitFlag(userId int64) {
	client := cache.GetClient()
	key := cache.GenerateUserUsageLimitCacheKey(userId)
	if err := client.Set(c, key, "", time.Hour*1).Err(); err != nil {
		logging.Warn(err)
	}
}
