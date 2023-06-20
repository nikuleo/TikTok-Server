package controller

import (
	"TikTokServer/pkg/tlog"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	tlog.Info("CommentAction")
}

func GetCommentList(c *gin.Context) {
	tlog.Info("CommentList")
}
