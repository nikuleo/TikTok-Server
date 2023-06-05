package controller

import (
	"TikTokServer/app/log"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(ctx *gin.Context) {
	// TODO:
	log.Info("FavoriteAction")
}

func GetFavoriteList(ctx *gin.Context) {
	log.Info("FavoriteList")
}
