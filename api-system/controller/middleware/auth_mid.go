package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/axiangcoding/antonstar-bot/entity/e"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
)

// CqhttpAuth 判断X-Self-ID是否和配置项相同，同时当X-Signature存在时，校验签名
func CqhttpAuth(selfQQ int64, secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		hSelfQQStr := c.GetHeader("X-Self-ID")
		hSignature := c.GetHeader("X-Signature")
		hSelfQQ, err := strconv.ParseInt(hSelfQQStr, 10, 64)
		if err != nil {
			app.Unauthorized(c, e.TokenNotValid, err)
			return
		}
		if selfQQ != hSelfQQ {
			app.Unauthorized(c, e.TokenNotValid)
			return
		}
		if hSignature != "" {
			if secret == "" {
				app.Unauthorized(c, e.TokenNotValid)
				return
			}
			body, err := io.ReadAll(c.Request.Body)
			if err != nil {
				app.Unauthorized(c, e.TokenNotValid, err)
				return
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			mac := hmac.New(sha1.New, []byte(secret))
			if _, err := mac.Write(body); err != nil {
				app.Unauthorized(c, e.TokenNotValid, err)
				return
			}
			if "sha1="+hex.EncodeToString(mac.Sum(nil)) != hSignature {
				app.Unauthorized(c, e.TokenNotValid)
				return
			}
		}
		c.Next()
	}

}
