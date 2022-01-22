package service

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/pkg/cache"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"time"
)

const tokenKeyPrefix = "Authorization#"

func getRefreshDuration() time.Duration {
	duration := conf.Config.App.Auth.RefreshDuration
	parseDuration, err := time.ParseDuration(duration)
	if err != nil {
		logging.Errorf("refresh_duration config valid, please check it again. ", err)
		return 5 * time.Minute
	}
	return parseDuration
}

func CacheToken(ctx *gin.Context, key string, token string) error {
	err := cache.GetRedis().Set(ctx, tokenKeyPrefix+key, token, getRefreshDuration()).Err()
	return err
}

func CacheTokenExist(ctx *gin.Context, key string) (bool, error) {
	result, err := cache.GetRedis().Exists(ctx, tokenKeyPrefix+key).Result()
	if err != nil {
		return false, err
	}
	if result > 0 {
		return true, nil
	}
	return false, nil
}

func RefreshTokenTTL(ctx *gin.Context, key string) {
	err := cache.GetRedis().Expire(ctx, tokenKeyPrefix+key, getRefreshDuration()).Err()
	if err != nil {
		logging.Error(err)
	}
}

func GetCachedToken(ctx *gin.Context, key string) (string, error) {
	result, err := cache.GetRedis().Get(ctx, tokenKeyPrefix+key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return result, nil
}

func DeleteCachedToken(ctx *gin.Context, key string) (int64, error) {
	result, err := cache.GetRedis().Del(ctx, tokenKeyPrefix+key).Result()
	return result, err
}
