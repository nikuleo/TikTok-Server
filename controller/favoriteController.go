package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(ctx *gin.Context) {
	tlog.Info("FavoriteAction")
	var err error
	videoID, err := strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	actionType, err := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	userID, _ := ctx.Get("userID")
	authID := userID.(int64)

	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}

	resp, err := service.FavoriteAction(videoID, authID, actionType)

	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}

// authID 自己的鉴权 id， userID 目标用户 id
func GetFavoriteList(ctx *gin.Context) {
	tlog.Info("FavoriteList")

	var err error

	aID, _ := ctx.Get("userID")
	authID := aID.(int64)
	userID, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)

	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}

	resp, err := service.GetFavoriteList(authID, userID)
	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}
