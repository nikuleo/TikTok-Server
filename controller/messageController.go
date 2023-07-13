package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MessageAction(c *gin.Context) {
	tlog.Info("MessageAction")
	toUserID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		response.Fail(c, errorcode.ErrHttpBind, nil)
		return
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil {
		response.Fail(c, errorcode.ErrHttpBind, nil)
		return
	}

	content := c.Query("content")
	aID, _ := c.Get("userID")
	authID := aID.(int64)

	resp, err := service.MessageAction(authID, toUserID, actionType, content)

	if err != nil {
		response.Fail(c, err, nil)
		return
	}

	response.Success(c, errorcode.HttpSuccess, resp)
}

func GetMessageList(ctx *gin.Context) {
	tlog.Info("GetChatMessages")

	toUserID, err := strconv.ParseInt(ctx.Query("to_user_id"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}

	aID, _ := ctx.Get("userID")
	authID := aID.(int64)

	resp, err := service.GetMessageList(authID, toUserID)
	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}
	response.Success(ctx, errorcode.HttpSuccess, resp)
}

func GetFriendList(ctx *gin.Context) {
	tlog.Info("GetFriendList")

	userID, err := strconv.ParseInt(ctx.Query("user_id"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}

	aID, _ := ctx.Get("userID")
	authID := aID.(int64)
	if userID != authID {
		response.Fail(ctx, errorcode.ErrHttpTokenInvalid, nil)
		return
	}

	resp, err := service.GetFriendList(authID)
	tlog.Infof("GetFriendList resp: %v", resp)
	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}
