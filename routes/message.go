package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addMessageRoutes(rg *gin.RouterGroup) {
	message := rg.Group("/message")
	message.Use(middleware.JwtAuthMiddleware())
	message.POST("/action/", controller.MessageAction)
	message.GET("/chat/", controller.GetChatMessages)
}
