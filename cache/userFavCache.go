package cache

import (
	"TikTokServer/pkg/util"
	"strconv"
	"time"
)

// 使用 set 存储 key-"user:id" 点赞视频 videoID-"id"
func SetUserFavVideos(userID int64, favList []int64) error {
	key := "user:" + strconv.FormatInt(userID, 10)

	values := make([]string, len(favList))
	for i, videoid := range favList {
		values[i] = strconv.FormatInt(videoid, 10)
	}
	err := RdbUserFavorite.SAdd(Ctx, key, values).Err()
	if err != nil {
		return err
	}
	_, err = RdbUserFavorite.Expire(Ctx, key, time.Duration(EXPIRE)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetUserFavVideos(userID int64) ([]int64, error) {
	key := "user:" + strconv.FormatInt(userID, 10)
	values, err := RdbUserFavorite.SMembers(Ctx, key).Result()
	if err != nil {
		return nil, err
	}
	viedoIDSlice, err := util.ConvtStrSliceToInt64Slice(values)
	if err != nil {
		return nil, err
	}
	return viedoIDSlice, nil
}

func DelUserFavVideos(userID int64) error {
	key := "user:" + strconv.FormatInt(userID, 10)
	err := RdbUserFavorite.Del(Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
