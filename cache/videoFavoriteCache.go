package cache

import (
	"TikTokServer/pkg/util"
	"strconv"
)

// 使用 set 存储 key-"video:id" 点赞用户 userID-"id"，过期时间前持久化到数据库
func SetVideoFavoriteUserToCache(videoID, userID int64) error {
	key := "video:" + strconv.FormatInt(videoID, 10)

	err := RdbVideoFavorite.SAdd(Ctx, key, strconv.FormatInt(userID, 10)).Err()
	if err != nil {
		return err
	}

	// 该缓存不设过期时间
	// ttl, err := RdbVideoFavorite.TTL(Ctx, key).Result()
	// if err != nil {
	// 	return err
	// }

	// // 如果没设置过期时间，设置过期时间
	// if ttl == -1 {
	// 	_, err = RdbVideoFavorite.Expire(Ctx, key, time.Duration(EXPIRE)*time.Second).Result()
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func DelVideoFavoriteUserInCache(videoID, userID int64) error {
	key := "video:" + strconv.FormatInt(videoID, 10)

	err := RdbVideoFavorite.SRem(Ctx, key, strconv.FormatInt(userID, 10)).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetVideoFavoriteUserFromCache(videoID int64) ([]int64, error) {
	key := "video:" + strconv.FormatInt(videoID, 10)

	userIDs, err := RdbVideoFavorite.SMembers(Ctx, key).Result()
	if err != nil {
		return nil, err
	}

	userIDsInt64, err := util.ConvtStrSliceToInt64Slice(userIDs)
	if err != nil {
		return nil, err
	}

	return userIDsInt64, nil
}
