package cache

import (
	"context"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/go-redis/redis/v8"
)

var _rdb *redis.Client

func Setup() {
	_rdb = initRedis()
	err := _rdb.Ping(context.Background()).Err()
	if err != nil {
		logging.L().Fatal("redis connect failed", logging.Error(err))
	}
}

func initRedis() *redis.Client {
	source := settings.C().App.Data.Cache.Source
	opt, err := redis.ParseURL(source)
	if err != nil {
		logging.L().Fatal("parse redis url failed",
			logging.Error(err),
			logging.Any("source", source))
	}
	return redis.NewClient(opt)
}

func Client() *redis.Client {
	return _rdb
}
