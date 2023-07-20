package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/ossBucket"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// authID 为当前用户ID, userID 目标用户ID （场景： 我点开了目标用户的发布列表）
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
	videoUrl, err := ossBucket.UploadVideoToOss(fileName, savePath)
	if err != nil {
		return nil, err
	}

	coverPath, coverName, err := GetImageFile(savePath)
	if err != nil {
		return nil, err
	}
	// tlog.Debugf("coverPath: %s", coverPath)
	coverUrl, err := ossBucket.UploadCoverToOss(coverName, coverPath)

	if err != nil {
		return nil, err
	}

	err = model.CreateVideo(userID, videoUrl, coverUrl, title)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	return &message.DouyinPublishActionResponse{}, nil
}

func PackVideoList(videos []*model.Video, userID int64) []*message.Video {
	followList, err := getFollowUserIDs(userID)
	if err != nil {
		return nil
	}
	favList, err := getUserFavVideoIDList(userID)
	if err != nil {
		return nil
	}

	followMap := make(map[int64]struct{})
	favMap := make(map[int64]struct{})
	for _, v := range followList {
		followMap[v] = struct{}{}
	}
	for _, v := range favList {
		favMap[v] = struct{}{}
	}

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
		if _, ok := followMap[int64(v.Author.ID)]; ok {
			video.Author.IsFollow = true
		}
		if _, ok := favMap[int64(v.ID)]; ok {
			video.IsFavorite = true
		}
		videoList[i] = video
	}
	return videoList
}

// 视屏 FFmpeg 截取视频封面
func GetImageFile(videoPath string) (string, string, error) {
	temp := strings.Split(videoPath, "/")
	videoName := temp[len(temp)-1]
	b := []byte(videoName)
	coverName := string(b[:len(b)-3]) + "jpg"
	homePath := os.Getenv("HOME")
	coverPath := filepath.Join(homePath, "/tmp/tiktokserver/cover/", coverName)
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-ss", "1", "-f", "image2", "-t", "0.01", "-y", coverPath)
	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	return coverPath, coverName, nil
}
