package controller

import (
	"TikTokServer/pkg/log"

	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	log.Info("RelationAction")
}

func GetFollowList(c *gin.Context) {
	log.Info("GetFollowList")
}

func GetFollowerList(c *gin.Context) {
	log.Info("GetFollowerList")
}

func GetFriendList(c *gin.Context) {
	log.Info("GetFriendList")
}
