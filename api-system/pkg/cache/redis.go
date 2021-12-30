package cache

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"context"
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
	opt, err := redis.ParseURL(conf.Config.App.Data.Cache.Source)
	if err != nil {
		logging.Fatal(err)
	}
	return redis.NewClient(opt)
}

func GetRedis() *redis.Client {
	return rdb
}
