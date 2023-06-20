package controller

import (
	"TikTokServer/pkg/tlog"

	"github.com/gin-gonic/gin"
)

func Feed(ctx *gin.Context) {
	tlog.Info("feedController")
}
