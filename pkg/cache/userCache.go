package cache

import (
	message "TikTokServer/idl/gen"
	"strconv"
)

func SetUserInfo(userID int64, userInfo *message.User) error {
	key := "user:" + strconv.FormatInt(userID, 10)
	value := packUserInfo(userInfo)
	err := RdbUserInfo.HSet(Ctx, key, value).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(userID int64) (*message.User, error) {
	key := "user:" + strconv.FormatInt(userID, 10)
	value, err := RdbUserInfo.HGetAll(Ctx, key).Result()
	if err != nil {
		return nil, err
	}

	followCount, _ := strconv.ParseInt(value["follow_count"], 10, 64)
	followerCount, _ := strconv.ParseInt(value["follower_count"], 10, 64)
	totalFavorited, _ := strconv.ParseInt(value["total_favorited"], 10, 64)
	workCount, _ := strconv.ParseInt(value["work_count"], 10, 64)
	favoriteCount, _ := strconv.ParseInt(value["favorite_count"], 10, 64)

	return &message.User{
		Id:              userID,
		Name:            value["name"],
		FollowCount:     followCount,
		FollowerCount:   followerCount,
		Avatar:          value["avatar"],
		BackgroundImage: value["background_image"],
		Signature:       value["signature"],
		TotalFavorited:  totalFavorited,
		WorkCount:       workCount,
		FavoriteCount:   favoriteCount,
	}, nil
}
