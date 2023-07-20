package cache

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/pkg/util"
	"strconv"
	"time"
)

// 使用 set 存储 key-"user:id" 与对应关注者 value-"id"
func SetUserFollowing(userID int64, followList []*message.User) error {
	key := "user:" + strconv.FormatInt(userID, 10)

	values := make([]string, len(followList))
	for i, user := range followList {
		values[i] = strconv.FormatInt(user.Id, 10)
	}
	err := RdbUserFollowing.SAdd(Ctx, key, values).Err()
	if err != nil {
		return err
	}
	_, err = RdbUserFollowing.Expire(Ctx, key, time.Duration(EXPIRE)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetUserFollowing(userID int64) ([]int64, error) {
	key := "user:" + strconv.FormatInt(userID, 10)
	values, err := RdbUserFollowing.SMembers(Ctx, key).Result()
	if err != nil {
		return nil, err
	}
	followIDList, err := util.ConvtStrSliceToInt64Slice(values)
	if err != nil {
		return nil, err
	}
	return followIDList, nil
}

func DelUserFollowing(userID int64) error {
	key := "user:" + strconv.FormatInt(userID, 10)
	err := RdbUserFollowing.Del(Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
