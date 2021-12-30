package middleware

import (
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		end := time.Now()
		latencyTime := end.Sub(start)
		reqMethod := c.Request.Method
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logging.Infof("%s %s %s--> status=%d, latency_time=%s, ip=%s",
			reqMethod, path, raw, statusCode, latencyTime, clientIP)
	}
}
