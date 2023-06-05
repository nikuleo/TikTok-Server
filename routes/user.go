package routes

import (
	"TikTokServer/controller"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/", controller.GetUserInfo)
	user.POST("/register/", controller.UserRegister)
	user.POST("/login/", controller.UserLogin)
}
