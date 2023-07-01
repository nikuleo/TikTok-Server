package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/", middleware.JwtAuthMiddleware(), controller.GetUserInfo)
	user.POST("/register/", controller.UserRegister)
	user.POST("/login/", controller.UserLogin)
}
