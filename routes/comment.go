package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addCommentRoutes(rg *gin.RouterGroup) {
	comment := rg.Group("/comment")
	comment.Use(middleware.JwtAuthMiddleware())
	comment.POST("/action/", controller.CommentAction)
	comment.GET("/list/", controller.GetCommentList)
}
