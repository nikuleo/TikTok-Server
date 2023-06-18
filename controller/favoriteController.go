package controller

import (
	"TikTokServer/pkg/tlog"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(ctx *gin.Context) {
	// TODO:
	tlog.Info("FavoriteAction")
}

func GetFavoriteList(ctx *gin.Context) {
	tlog.Info("FavoriteList")
}
