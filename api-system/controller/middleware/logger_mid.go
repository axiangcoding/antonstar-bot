package middleware

import (
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latencyTime := end.Sub(start)
		reqMethod := c.Request.Method
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		fields := []zap.Field{
			logging.Any("method", reqMethod),
			logging.Any("path", path),
			logging.Any("raw", raw),
			logging.Any("statusCode", statusCode),
			logging.Any("latencyTime", latencyTime),
			logging.Any("clientIp", clientIP),
		}
		if latencyTime > time.Second*3 {
			logging.L().Warn("receive slow request", fields...)
		} else {
			logging.L().Info("receive request", fields...)
		}

	}
}
