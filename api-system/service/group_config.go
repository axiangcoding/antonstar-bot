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

func FindGroupConfig(groupId int64) (*table.QQGroupConfig, error) {
	db := data.GetDB()
	var find table.QQGroupConfig
	if err := db.Where(table.QQGroupConfig{GroupId: groupId}).Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func MustFindGroupConfig(groupId int64) *table.QQGroupConfig {
	config, err := FindGroupConfig(groupId)
	if err != nil {
		logging.Warn(err)
	}
	return config
}

func GetGroupConfigWithCondition(condition table.QQGroupConfig) ([]table.QQGroupConfig, error) {
	db := data.GetDB()
	var finds []table.QQGroupConfig
	if err := db.Where(condition).Find(&finds).Error; err != nil {
		return nil, err
	}
	return finds, nil
}

func SaveGroupConfig(gc table.QQGroupConfig) error {
	db := data.GetDB()
	if err := db.Save(&gc).Error; err != nil {
		return err
	}
	return nil
}

func ExistBiliRoomFlag(groupId int64, roomId int64) bool {
	client := cache.GetClient()
	key := cache.GenerateBiliRoomLivingCacheKey(groupId, roomId)
	if _, err := client.Get(c, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return false
		}
		logging.Warn(err)
		return false
	}
	return true
}

func MustPutBiliRoomFlag(groupId int64, roomId int64) {
	client := cache.GetClient()
	key := cache.GenerateBiliRoomLivingCacheKey(groupId, roomId)
	if err := client.Set(c, key, "", time.Minute*10).Err(); err != nil {
		logging.Warn(err)
	}
}
