package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addFavoriteRoutes(rg *gin.RouterGroup) {
	favorite := rg.Group("/favorite")

	favorite.POST("/action/", controller.FavoriteAction)
	favorite.GET("/list/", controller.GetFavoriteList)
}
