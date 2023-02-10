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

func MustSaveGameProfile(gameUser *table.GameUser) {
	if err := dal.Q.GameUser.Save(gameUser); err != nil {
		logging.L().Error("dal error", logging.Error(err))
	}
}

func MustUpdateGameProfile(nick string, user *table.GameUser) {
	if _, err := dal.GameUser.Where(dal.GameUser.Nick.Eq(nick)).Updates(user); err != nil {
		logging.L().Error("dal error", logging.Error(err))
	}
}

func CanBeRefresh(nick string) bool {
	client := cache.GetClient()
	key := cache.GenerateGameUserCacheKey(nick)
	if _, err := client.Get(context.Background(), key).Result(); err != nil {
		if errors.Is(err, redis.Nil) {
			return true
		}
		logging.L().Error("get cache error", logging.Error(err))
		return false
	}
	return false
}

func MustPutRefreshFlag(nick string) {
	client := cache.GetClient()
	key := cache.GenerateGameUserCacheKey(nick)
	if err := client.Set(context.Background(), key, "", time.Hour*24).Err(); err != nil {
		logging.L().Error("set cache error", logging.Error(err))
	}
}
