package middleware

import (
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/entity/e"
	"github.com/axiangcoding/ax-web/settings"
	"github.com/gin-gonic/gin"
)

func AppToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenInParam := c.Query("app_token")
		if settings.Config.Server.AppToken == tokenInParam {
			c.Next()
		} else {
			app.Unauthorized(c, e.TokenNotValid)
			return
		}
	}
}
