package controller

import (
	"TikTokServer/pkg/tlog"

	"github.com/gin-gonic/gin"
)

func PublishAction(ctx *gin.Context) {
	tlog.Info("publish action")
}

func GetPublishList(ctx *gin.Context) {
	tlog.Info("publish list")
}
