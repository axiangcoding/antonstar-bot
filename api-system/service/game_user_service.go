package service

import (
	"context"
	"errors"
	"github.com/axiangcoding/antonstar-bot/cache"
	"github.com/axiangcoding/antonstar-bot/data/dal"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/go-redis/redis/v8"
	"regexp"
	"time"
)

var c = context.Background()

func IsValidNickname(nick string) bool {
	matched, err := regexp.Match(`^[\w\s@]+$`, []byte(nick))
	if err != nil {
		return false
	}
	return matched
}

func FindGameProfile(nick string) (*table.GameUser, error) {
	take, err := dal.Q.GameUser.Where(dal.GameUser.Nick.Eq(nick)).Take()
	if err != nil {
		return nil, err
	}
	return take, err
}

func SaveGameProfile(gameUser table.GameUser) error {
	if err := dal.Q.GameUser.Save(&gameUser); err != nil {
		return err
	}
	return nil
}

func UpdateGameProfile(nick string, user table.GameUser) error {
	if _, err := dal.GameUser.Where(dal.GameUser.Nick.Eq(nick)).Updates(&user); err != nil {
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
