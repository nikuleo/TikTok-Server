package controller

import (
	"TikTokServer/app/log"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	log.Info("CommentAction")
}

func GetCommentList(c *gin.Context) {
	log.Info("CommentList")
}
