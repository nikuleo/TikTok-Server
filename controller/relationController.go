package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RelationAction(ctx *gin.Context) {
	tlog.Info("RelationAction")

	actionType, err := strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}
	toUserID, err := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}

	aID, _ := ctx.Get("userID")
	authID := aID.(int64)

	resp, err := service.RelationAction(authID, toUserID, actionType)
	tlog.Infof("relationAction resp: %v", resp)
	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}

func GetFollowList(c *gin.Context) {
	tlog.Info("GetFollowList")

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		response.Fail(c, errorcode.ErrHttpBind, nil)
		return
	}

	aID, _ := c.Get("userID")
	authID := aID.(int64)

	if userID != authID {
		response.Fail(c, errorcode.ErrHttpTokenInvalid, nil)
		return
	}

	resp, err := service.GetFollowList(authID)

	if err != nil {
		response.Fail(c, err, nil)
		return
	}

	response.Success(c, errorcode.HttpSuccess, resp)
}

func GetFollowerList(c *gin.Context) {
	tlog.Info("GetFollowerList")

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		response.Fail(c, errorcode.ErrHttpBind, nil)
		return
	}

	aID, _ := c.Get("userID")
	authID := aID.(int64)
	if userID != authID {
		response.Fail(c, errorcode.ErrHttpTokenInvalid, nil)
		return
	}

	resp, err := service.GetFollowerList(authID)

	if err != nil {
		response.Fail(c, err, nil)
		return
	}

	response.Success(c, errorcode.HttpSuccess, resp)
}
