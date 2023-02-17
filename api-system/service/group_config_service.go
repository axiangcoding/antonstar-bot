package service

import (
	"context"
	"errors"
	"github.com/axiangcoding/antonstar-bot/cache"
	"github.com/axiangcoding/antonstar-bot/data/dal"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/go-redis/redis/v8"
	"time"
)

func FindGroupConfig(groupId int64) (*table.QQGroupConfig, error) {
	take, err := dal.Q.QQGroupConfig.Where(dal.QQGroupConfig.GroupId.Eq(groupId)).Take()
	if err != nil {
		return nil, err
	}
	return take, nil
}

func MustFindGroupConfig(groupId int64) *table.QQGroupConfig {
	config, err := FindGroupConfig(groupId)
	if err != nil {
		logging.L().Warn("dal failed", logging.Error(err))
	}
	return config
}

func GetEnableCheckBiliRoomGroupConfig(enableCheckBiliBiliRoom bool) ([]*table.QQGroupConfig, error) {
	find, err := dal.QQGroupConfig.Where(dal.QQGroupConfig.EnableCheckBiliRoom.Is(enableCheckBiliBiliRoom)).Find()
	return find, err
}

func SaveGroupConfig(gc table.QQGroupConfig) error {
	if err := dal.QQGroupConfig.Save(&gc); err != nil {
		return err
	}
	return nil
}

func MustSaveGroupConfig(gc *table.QQGroupConfig) {
	if err := dal.QQGroupConfig.Save(gc); err != nil {
		logging.L().Error("dal failed", logging.Error(err))
	}
}

func ExistBiliRoomFlag(groupId int64, roomId int64) bool {
	client := cache.Client()
	key := cache.GenerateBiliRoomLivingCacheKey(groupId, roomId)
	if _, err := client.Get(context.Background(), key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return false
		}
		logging.L().Warn("get cache failed", logging.Error(err))
		return false
	}
	return true
}

func ExistGroupUsageLimitFlag(groupId int64) bool {
	client := cache.Client()
	key := cache.GenerateGroupUsageLimitCacheKey(groupId)
	if _, err := client.Get(context.Background(), key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return false
		}
		logging.L().Warn("get cache failed", logging.Error(err))
		return false
	}
	return true
}

func MustPutBiliRoomFlag(groupId int64, roomId int64) {
	client := cache.Client()
	key := cache.GenerateBiliRoomLivingCacheKey(groupId, roomId)
	if err := client.Set(context.Background(), key, "", time.Minute*10).Err(); err != nil {
		logging.L().Warn("set cache failed", logging.Error(err))
	}
}

func MustPutGroupUsageLimitFlag(groupId int64) {
	client := cache.Client()
	key := cache.GenerateGroupUsageLimitCacheKey(groupId)
	if err := client.Set(context.Background(), key, "", time.Hour*1).Err(); err != nil {
		logging.L().Warn("set cache failed", logging.Error(err))
	}
}

func CheckGroupTodayQueryLimit(groupId int64) (bool, int, int) {
	config := MustFindGroupConfig(groupId)
	if config == nil {
		return true, 0, 0
	}
	return config.TodayQueryCount >= config.OneDayQueryLimit, config.TodayQueryCount, config.OneDayQueryLimit
}

func MustAddGroupConfigTodayQueryCount(groupId int64, count int) {
	config := MustFindGroupConfig(groupId)
	config.TodayQueryCount += count
	MustSaveGroupConfig(config)
}

func MustAddGroupConfigTotalQueryCount(groupId int64, count int) {
	config := MustFindGroupConfig(groupId)
	config.TotalQueryCount += count
	MustSaveGroupConfig(config)
}

func CheckGroupTodayUsageLimit(groupId int64) (bool, int, int) {
	config := MustFindGroupConfig(groupId)
	if config == nil {
		return true, 0, 0
	}
	return config.TodayUsageCount >= config.OneDayUsageLimit, config.TodayUsageCount, config.OneDayUsageLimit
}

func MustAddGroupConfigTodayUsageCount(groupId int64, count int) {
	config := MustFindGroupConfig(groupId)
	config.TodayUsageCount += count
	MustSaveGroupConfig(config)
}

func MustAddGroupConfigTotalUsageCount(groupId int64, count int) {
	config := MustFindGroupConfig(groupId)
	config.TotalUsageCount += count
	MustSaveGroupConfig(config)
}

func ResetAllGroupConfigTodayCount() error {
	quc := dal.QQGroupConfig
	if _, err := dal.QQGroupConfig.
		Select(quc.TodayQueryCount, quc.TodayUsageCount).
		Where(quc.ID.IsNotNull()).
		Updates(table.QQGroupConfig{TodayQueryCount: 0, TodayUsageCount: 0}); err != nil {
		return err
	}
	return nil
}
