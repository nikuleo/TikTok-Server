package controller

import (
	"TikTokServer/pkg/auth"
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/response"
	"TikTokServer/pkg/tlog"
	"TikTokServer/pkg/util"
	"TikTokServer/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 可以不登录访问，如果有 token 手动鉴权
func Feed(ctx *gin.Context) {
	tlog.Info("feedController")
	var userID int64 = -1
	var err error
	var latestTime int64
	token := ctx.Query("token")
	latestTime, err = strconv.ParseInt(ctx.Query("latest_time"), 10, 64)

	if err != nil || latestTime == int64(0) {
		latestTime = util.GetCurrentTime()
	}
	if token != "" {
		userID, err = auth.GetUserIDByToken(token)
		if err != nil {
			errCode := errorcode.ErrHttpTokenInvalid
			errCode.SetError(err)
			response.Fail(ctx, errCode, nil)
			return
		}
	}

	resp, err := service.GetFeedList(latestTime, userID)

	if err != nil {
		response.Fail(ctx, err, nil)
		return
	}
	tlog.Infof("feedList: %v", resp)
	response.Success(ctx, errorcode.HttpSuccess, resp)
}
