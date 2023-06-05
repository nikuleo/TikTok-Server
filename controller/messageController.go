package controller

import (
	"TikTokServer/app/log"

	"github.com/gin-gonic/gin"
)

func MessageAction(c *gin.Context) {
	log.Info("MessageAction")
}

func GetChatMessages(c *gin.Context) {
	log.Info("GetChatMessages")
}
