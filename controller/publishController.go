package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PublishAction(ctx *gin.Context) {
	tlog.Info("publish action")
}

func GetPublishList(ctx *gin.Context) {
	tlog.Info("publish list")
	userID := ctx.Query("user_id")
	authID, _ := ctx.Get("userID")
	aID := authID.(int64)

	if userID == "" {
		userID = strconv.FormatInt(aID, 10)
	}
	uID, _ := strconv.Atoi(userID)
	resp, err := service.PublishList(aID, int64(uID))

	if err != nil {
		tlog.Errorf(err.Error(), resp)
		response.Fail(ctx, err, nil)
		return
	}
	tlog.Infof("publishList: %v", resp)
	response.Success(ctx, errorcode.HttpSuccess, resp)
}
