package service

import (
	"TikTokServer/cache"
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/tlog"
	"TikTokServer/rabbitmq"
	"fmt"
)

func FavoriteAction(videoID, userID, actionType int64) (*message.DouyinFavoriteActionResponse, error) {
	var err error
	msg := fmt.Sprintf("%d %d", userID, videoID)

	// 使用消息队列优化
	if actionType == 1 {
		// err = model.Favorite(userID, videoID)
		err = cache.SetVideoFavoriteUserToCache(videoID, userID)
		if err != nil {
			return nil, err
		}
		err = rabbitmq.SendFavoriteMsg(msg)
	}

	if actionType == 2 {
		// err = model.DisFavorite(userID, videoID)
		err = cache.DelVideoFavoriteUserInCache(videoID, userID)
		if err != nil {
			return nil, err
		}
		err = rabbitmq.SendDisFavoriteMsg(msg)
	}

	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	resp := &message.DouyinFavoriteActionResponse{}

	return resp, nil
}

func GetFavoriteList(authID, userID int64) (*message.DouyinFavoriteListResponse, error) {
	videoList, err := model.GetFavoriteList(userID)

	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	resp := &message.DouyinFavoriteListResponse{
		VideoList: PackVideoList(videoList, authID),
	}

	tlog.Infof("resp: %v", resp)

	return resp, nil
}

func getUserFavVideoIDList(userID int64) ([]int64, error) {
	videoIDs, err := cache.GetUserFavVideos(userID)
	if err != nil {
		return nil, err
	}

	if len(videoIDs) != 0 {
		return videoIDs, nil
	}
	videoList, err := model.GetFavoriteList(userID)
	if err != nil {
		return nil, err
	}
	videoIDs = make([]int64, len(videoList))
	for i, video := range videoList {
		videoIDs[i] = int64(video.ID)
	}
	if len(videoIDs) != 0 {
		err = cache.SetUserFavVideos(userID, videoIDs)
		if err != nil {
			return nil, err
		}
	}
	return videoIDs, nil
}
