package middleware

import (
	"TikTokServer/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLog() gin.HandlerFunc {
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
		log.Info(
			"GIN",
			log.String("IP", clientIP),
			log.String("path", path),
			log.String("method", method),
			log.Int("status", status),
			log.String("query", query),
			log.String("userAgent", userAgent),
			log.String("error", ctx.Errors.ByType(gin.ErrorTypePrivate).String()),
			log.Duration("duration", duration),
		)
	}
}
