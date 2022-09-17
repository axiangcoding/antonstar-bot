package service

import (
	"context"
	"errors"
	"github.com/axiangcoding/ax-web/cache"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/go-redis/redis/v8"
	"regexp"
	"time"
)

var c = context.Background()

func IsValidNickname(nick string) bool {
	matched, err := regexp.Match(`^\w+$`, []byte(nick))
	if err != nil {
		return false
	}
	return matched
}

func FindGameProfile(nick string) (*table.GameUser, error) {
	db := data.GetDB()
	var find table.GameUser
	if err := db.Where(table.GameUser{Nick: nick}).Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func SaveGameProfile(gameUser table.GameUser) error {
	db := data.GetDB()
	if err := db.Save(&gameUser).Error; err != nil {
		return err
	}
	return nil
}

func UpdateGameProfile(nick string, user table.GameUser) error {
	db := data.GetDB()
	if err := db.Where(table.GameUser{Nick: nick}).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func CanBeRefresh(nick string) bool {
	client := cache.GetClient()
	key := cache.GenerateGameUserCacheKey(nick)
	if _, err := client.Get(c, key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return true
		}
		logging.Warn(err)
		return false
	}
	return false
}

func MustPutRefreshFlag(nick string) {
	client := cache.GetClient()
	key := cache.GenerateGameUserCacheKey(nick)
	if err := client.Set(c, key, "", time.Hour*24).Err(); err != nil {
		logging.Warn(err)
	}
}
