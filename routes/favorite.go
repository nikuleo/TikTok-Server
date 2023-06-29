package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addFavoriteRoutes(rg *gin.RouterGroup) {
	favorite := rg.Group("/favorite")
	favorite.Use(middleware.JwtAuthMiddleware())
	favorite.POST("/action/", controller.FavoriteAction)
	favorite.GET("/list/", controller.GetFavoriteList)
}
