package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addPublishRoutes(rg *gin.RouterGroup) {
	publish := rg.Group("/publish")

	publish.POST("/action/", controller.PublishAction)
	publish.GET("/list/", controller.GetPublishList)
}
