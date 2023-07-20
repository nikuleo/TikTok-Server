package cache

import (
	message "TikTokServer/idl/gen"
	"encoding/json"
	"strconv"
	"time"
)

// 修改方案序列化后使用 string 存储
func SetUserInfo(userID int64, userInfo *message.User) error {
	key := "user:" + strconv.FormatInt(userID, 10)
	value, err := json.Marshal(userInfo)
	if err != nil {
		return err
	}
	err = RdbUserInfo.Set(Ctx, key, value, time.Duration(EXPIRE)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(userID int64) (*message.User, error) {
	key := "user:" + strconv.FormatInt(userID, 10)
	value, err := RdbUserInfo.Get(Ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}

	var userInfo message.User
	err = json.Unmarshal([]byte(value), &userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func DelUserInfo(userID int64) error {
	key := "user:" + strconv.FormatInt(userID, 10)
	err := RdbUserInfo.Del(Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
