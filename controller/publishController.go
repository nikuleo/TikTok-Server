package controller

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/service"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PublishAction(ctx *gin.Context) {
	tlog.Info("publish action")
	userID, _ := ctx.Get("userID")
	aID := userID.(int64)
	title := ctx.PostForm("title")
	videoData, err := ctx.FormFile("data")
	if err != nil {
		errCode := errorcode.ErrHttpBind
		errCode.SetError(err)
		response.Fail(ctx, errCode, nil)
		return
	}
	fileName := filepath.Base(videoData.Filename)

	// rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	currentTime := time.Now().UnixNano()

	fileName = fmt.Sprintf("%d_%d_%s", r.Intn(100000), currentTime, fileName)

	homePath := os.Getenv("HOME")
	savePath := filepath.Join(homePath, "/tmp/tiktokserver/video/", fileName)

	if err := ctx.SaveUploadedFile(videoData, savePath); err != nil {
		response.Fail(ctx, err, nil)
		return
	}

	resp, err := service.PublishAction(aID, title, fileName, savePath)

	if err != nil {
		tlog.Errorf(err.Error(), resp)
		response.Fail(ctx, err, nil)
		return
	}

	response.Success(ctx, errorcode.HttpSuccess, resp)
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
