package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addFeedRoutes(rg *gin.RouterGroup) {
	feed := rg.Group("/feed")
	feed.GET("/", controller.Feed)
}
