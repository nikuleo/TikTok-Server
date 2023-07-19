package cache

import message "TikTokServer/idl/gen"

type UserInfo struct {
	Id              int64  `redis:"id"`
	Name            string `redis:"name"`
	FollowCount     int64  `redis:"follow_count"`
	FollowerCount   int64  `redis:"follower_count"`
	Avatar          string `redis:"avatar"`
	BackgroundImage string `redis:"background_image"`
	Signature       string `redis:"signature"`
	TotalFavorited  int64  `redis:"total_favorited"`
	WorkCount       int64  `redis:"work_count"`
	FavoriteCount   int64  `redis:"favorite_count"`
}

func packUserInfo(user *message.User) UserInfo {
	return UserInfo{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}
