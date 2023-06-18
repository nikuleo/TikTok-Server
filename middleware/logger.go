package middleware

import (
	"TikTokServer/pkg/tlog"
	"time"

	"github.com/gin-gonic/gin"
)

func Gintlog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		status := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		ctx.Next()
		userAgent := ctx.Request.UserAgent()
		duration := time.Since(start)
		tlog.Info(
			"GIN",
			tlog.String("IP", clientIP),
			tlog.String("path", path),
			tlog.String("method", method),
			tlog.Int("status", status),
			tlog.String("query", query),
			tlog.String("userAgent", userAgent),
			tlog.String("error", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			tlog.Duration("duration", duration),
		)
	}
}
