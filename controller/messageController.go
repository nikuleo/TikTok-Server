package controller

import (
	"TikTokServer/pkg/tlog"

	"github.com/gin-gonic/gin"
)

func MessageAction(c *gin.Context) {
	tlog.Info("MessageAction")
}

func GetChatMessages(c *gin.Context) {
	tlog.Info("GetChatMessages")
}
