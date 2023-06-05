package controller

import (
	"TikTokServer/app/log"

	"github.com/gin-gonic/gin"
)

func PublishAction(ctx *gin.Context) {
	log.Info("publish action")
}

func GetPublishList(ctx *gin.Context) {
	log.Info("publish list")
}
