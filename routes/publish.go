package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addPublishRoutes(rg *gin.RouterGroup) {
	publish := rg.Group("/publish")

	publish.POST("/action/", middleware.JwtAuthMiddleware(), controller.PublishAction)
	publish.GET("/list/", middleware.JwtAuthMiddleware(), controller.GetPublishList)
}
