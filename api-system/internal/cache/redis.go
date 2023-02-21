package cache

import (
	"context"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/go-redis/redis/v8"
)

var _rdb *redis.Client

func InitRedis(source string) {
	opt, err := redis.ParseURL(source)
	if err != nil {
		logging.L().Fatal("parse redis url failed",
			logging.Error(err),
			logging.Any("source", source))
	}
	_rdb = redis.NewClient(opt)
	err = _rdb.Ping(context.Background()).Err()
	if err != nil {
		logging.L().Fatal("redis connect failed", logging.Error(err))
	}
}

func Client() *redis.Client {
	return _rdb
}
