package cache

import (
	message "TikTokServer/idl/gen"
	"encoding/json"
	"strconv"
	"time"
)

// 修改方案序列化后使用 string 存储（每条评论序列大概 1 KB）
func SetVideoCommentToCache(videoID int64, comments []*message.Comment) error {
	key := "video:" + strconv.FormatInt(videoID, 10)
	value, err := json.Marshal(comments)
	if err != nil {
		return err
	}
	err = RdbVideoComment.Set(Ctx, key, value, time.Duration(EXPIRE)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetVideoCommentFromCache(videoID int64) ([]*message.Comment, error) {
	key := "video:" + strconv.FormatInt(videoID, 10)
	value, err := RdbVideoComment.Get(Ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}

	var comments []*message.Comment
	err = json.Unmarshal([]byte(value), &comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func DelVideoCommentCache(videoID int64) error {
	key := "video:" + strconv.FormatInt(videoID, 10)
	err := RdbVideoComment.Del(Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
