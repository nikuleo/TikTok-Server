package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addCommentRoutes(rg *gin.RouterGroup) {
	comment := rg.Group("/comment")
	comment.POST("/action/", controller.CommentAction)
	comment.GET("/list/", controller.GetCommentList)
}
