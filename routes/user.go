package routes

import (
	"TikTokServer/controller"
	"TikTokServer/middleware"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/", controller.GetUserInfo, middleware.JwtAuthMiddleware())
	user.POST("/register/", controller.UserRegister)
	user.POST("/login/", controller.UserLogin)
}
