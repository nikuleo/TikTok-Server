package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/tlog"
)

func FavoriteAction(videoID, userID, actionType int64) (*message.DouyinFavoriteActionResponse, error) {
	var err error
	if actionType == 1 {
		err = model.Favorite(userID, videoID)
	}

	if actionType == 2 {
		err = model.DisFavorite(userID, videoID)
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
