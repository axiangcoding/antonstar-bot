package middleware

import (
	"github.com/axiangcoding/ax-web/logging"
	"github.com/gin-gonic/gin"
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
		template := "%s %s %s --> status=%d, latency_time=%s, ip=%s"
		if latencyTime > time.Second*5 {
			logging.Warnf(template+" [slow response]", reqMethod, path, raw, statusCode, latencyTime, clientIP)
		} else {
			logging.Infof(template,
				reqMethod, path, raw, statusCode, latencyTime, clientIP)
		}

	}
}
