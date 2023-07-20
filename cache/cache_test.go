package cache

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/pkg/tlog"
	"fmt"
	"strconv"
	"testing"
)

func init() {
	tlog.InitLog()
	InitRedis()
}

func TestUserInfoCache(t *testing.T) {
	userInfo := &message.User{
		Id:              1,
		Name:            "niku",
		FollowCount:     1,
		FollowerCount:   1,
		Avatar:          "miku",
		BackgroundImage: "bg",
		Signature:       "嘿嘿嘿",
		TotalFavorited:  1,
		WorkCount:       1,
		FavoriteCount:   1,
	}

	err := SetUserInfo(1, userInfo)
	if err != nil {
		t.Error(err)
	}

	getUserInfo, err := GetUserInfo(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("getUserInfo: ", getUserInfo)

	nilUserInfo, err := GetUserInfo(2)
	if err != nil {
		fmt.Println("err:", err)
		t.Error(err)
	}
	fmt.Println("nilUserInfo: ", nilUserInfo)
}

func TestUserFollowingCache(t *testing.T) {
	userList := make([]*message.User, 10)
	for i := range userList {
		userList[i] = &message.User{
			Id:              int64(i),
			Name:            "niku" + strconv.Itoa(i),
			FollowCount:     1,
			FollowerCount:   1,
			Avatar:          "miku",
			BackgroundImage: "bg",
			Signature:       "嘿嘿嘿",
			TotalFavorited:  1,
			WorkCount:       1,
			FavoriteCount:   1,
		}
	}

	err := SetUserFollowing(1, userList)
	if err != nil {
		t.Error(err)
	}

	userFollowing, err := GetUserFollowing(1)
	if err != nil {
		fmt.Println("err:", err)
		t.Error(err)
	}
	fmt.Println("userFollowing: ", userFollowing)

	err = DelUserFollowing(1)
	if err != nil {
		t.Error(err)
	}
	userFollowing, err = GetUserFollowing(1)
	if err != nil {
		fmt.Println("err:", err)
		t.Error(err)
	}
	fmt.Println("Del userFollowing: ", userFollowing)
}
