package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {

	userName := ctx.Query("username")
	password := ctx.Query("password")

	resp, err := service.UserRegister(userName, password)

	if err != nil {
		tlog.Errorf(err.Error(), resp)
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)

}

func UserLogin(ctx *gin.Context) {

	tlog.Info("Usertlogin")

	userName := ctx.Query("username")
	password := ctx.Query("password")

	resp, err := service.UserLogin(userName, password)

	if err != nil {
		tlog.Errorf(err.Error(), resp)
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}

func GetUserInfo(ctx *gin.Context) {

	userID := ctx.Query("user_id")
	authID, _ := ctx.Get("userID")
	tlog.Debugf("userID: %v", userID)
	tlog.Debugf("authID: %v, %t", authID, authID)
	aID := authID.(int64)

	if strconv.FormatInt(aID, 10) != userID {
		response.Fail(ctx, errorcode.ErrHttpTokenInvalid, nil)
	}

	resp, err := service.GetUserInfo(aID)

	if err != nil {
		tlog.Errorf(err.Error(), resp)
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}
