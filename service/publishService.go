package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
)

func PublishList(authID, userID int64) (*message.DouyinPublishListResponse, error) {

	videos, err := model.GetVideoListByUserID(userID)

	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	resp := &message.DouyinPublishListResponse{
		VideoList: PackVideoList(videos, authID),
	}

	return resp, nil
}

func PublishAction(userID int64, title, fileName, savePath string) (*message.DouyinPublishActionResponse, error) {
	//TODO:
	// 上传视频到云存储
	// filePath, err := ossBucket.UploadFileToOss(fileName, savePath, "video/")

	return nil, nil
}

func PackVideoList(videos []*model.Video, userID int64) []*message.Video {
	//TODO: follow list & fav list, 写完关注接口与点赞后修改

	videoList := make([]*message.Video, len(videos))
	for i, v := range videos {
		video := &message.Video{
			Id:            int64(v.ID),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Author:        PackUserInfo(&v.Author),
			Title:         v.Title,
		}
		videoList[i] = video
	}
	return videoList
}
