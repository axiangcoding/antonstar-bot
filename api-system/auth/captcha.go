package auth

import (
	"context"
	"github.com/axiangcoding/ax-web/cache"
	"github.com/dchest/captcha"
	"time"
)

var c = context.Background()

func SetupCaptcha() {
	redisStore = &RedisStore{ExpireTime: time.Minute * 3}
	captcha.SetCustomStore(redisStore)
}

var captchaPrefix = "Captcha#"

// RedisStore 由于服务器是负载均衡，所以需要用redis来替代默认的store 如果是单台服务器，不需要redis来管理
var redisStore *RedisStore

type RedisStore struct {
	ExpireTime time.Duration // 过期时间
}

func (rs *RedisStore) Set(id string, digits []byte) {
	cache.GetClient().Set(c, captchaPrefix+id, string(digits), rs.ExpireTime)
}

func (rs *RedisStore) Get(id string, clear bool) (digits []byte) {
	result, _ := cache.GetClient().Get(c, captchaPrefix+id).Bytes()
	if clear {
		cache.GetClient().Del(c, captchaPrefix+id)
	}
	return result
}
