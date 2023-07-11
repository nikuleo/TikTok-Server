package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentAction(ctx *gin.Context) {
	tlog.Info("CommentAction")
	var (
		err         error
		videoID     int64
		actionType  int64
		commentText string = ""
		commentID   int64  = 0
	)
	videoID, err = strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}

	actionType, err = strconv.ParseInt(ctx.Query("action_type"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}
	userID, _ := ctx.Get("userID")
	authID := userID.(int64)

	if actionType == 1 {
		commentText = ctx.Query("comment_text")
	} else {
		commentID, err = strconv.ParseInt(ctx.Query("comment_id"), 10, 64)
		if err != nil {
			response.Fail(ctx, errorcode.ErrHttpBind, nil)
			return
		}
	}

	resp, err := service.CommentAction(authID, videoID, actionType, commentText, commentID)
	tlog.Infof("CommentAction resp: %v", resp)
	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}
	response.Success(ctx, errorcode.HttpSuccess, resp)

}

func GetCommentList(ctx *gin.Context) {

	var (
		err     error
		videoID int64
	)

	videoID, err = strconv.ParseInt(ctx.Query("video_id"), 10, 64)
	if err != nil {
		response.Fail(ctx, errorcode.ErrHttpBind, nil)
		return
	}
	userID, _ := ctx.Get("userID")
	authID := userID.(int64)

	resp, err := service.CommentList(authID, videoID)

	tlog.Infof("CommentList resp: %v", resp)

	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
}
