package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/util"
)

const (
	LIMIT = 30 // 返回最大视频数
)

func GetFeedList(latestTime, userID int64) (*message.DouyinFeedResponse, error) {
	videoList, err := model.GetVideoListByTime(latestTime, LIMIT)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	resp := &message.DouyinFeedResponse{
		VideoList: PackVideoList(videoList, userID),
	}

	nextTime := util.GetCurrentTime()
	if len(videoList) == LIMIT {
		nextTime = videoList[len(videoList)-1].PublishTime
	}
	resp.NextTime = nextTime
	return resp, nil
}
