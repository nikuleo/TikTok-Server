package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addPublishRoutes(rg *gin.RouterGroup) {
	publish := rg.Group("/publish")

	publish.POST("/action/", controller.PublishAction, middleware.JwtAuthMiddleware())
	publish.GET("/list/", controller.GetPublishList)
}
