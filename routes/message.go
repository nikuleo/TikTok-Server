package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addMessageRoutes(rg *gin.RouterGroup) {
	message := rg.Group("/message")
	message.POST("/action/", controller.MessageAction)
	message.GET("/chat/", controller.GetChatMessages)
}
