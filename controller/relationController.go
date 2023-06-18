package controller

import (
	"TikTokServer/pkg/tlog"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	tlog.Info("RelationAction")
}

func GetFollowList(c *gin.Context) {
	tlog.Info("GetFollowList")
}

func GetFollowerList(c *gin.Context) {
	tlog.Info("GetFollowerList")
}

func GetFriendList(c *gin.Context) {
	tlog.Info("GetFriendList")
}
