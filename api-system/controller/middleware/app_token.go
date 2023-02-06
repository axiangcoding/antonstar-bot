package middleware

import (
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/axiangcoding/antonstar-bot/entity/e"
	"github.com/axiangcoding/antonstar-bot/settings"
	"github.com/gin-gonic/gin"
)

func AppToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenInParam := c.Query("app_token")
		if settings.Config.Auth.AppToken == tokenInParam {
			c.Next()
		} else {
			app.Unauthorized(c, e.TokenNotValid)
			return
		}
	}
}
