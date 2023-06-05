package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addRelationRoutes(rg *gin.RouterGroup) {

	relation := rg.Group("/relation")

	relation.POST("/action/", controller.RelationAction)
	relation.GET("follow/list/", controller.GetFollowList)
	relation.GET("follower/list/", controller.GetFollowerList)
	relation.GET("friend/list/", controller.GetFriendList)
}
