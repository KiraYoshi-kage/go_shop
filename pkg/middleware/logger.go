package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		latency := end.Sub(start)

		// 打印日志
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// 这里先用gin默认的日志，后续可以替换成zap
		gin.DefaultWriter.Write([]byte(
			"method=" + c.Request.Method + " " +
				"path=" + path + " " +
				"status=" + string(c.Writer.Status()) + " " +
				"latency=" + latency.String() + "\n"))
	}
}
