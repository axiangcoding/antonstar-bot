package cache

import (
	"context"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Setup() {
	rdb = initRedis()
	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		logging.Fatal(err)
	}
}

func initRedis() *redis.Client {
	opt, err := redis.ParseURL(settings.Config.App.Data.Cache.Source)
	if err != nil {
		logging.Fatal(err)
	}
	return redis.NewClient(opt)
}

func GetRedis() *redis.Client {
	return rdb
}
