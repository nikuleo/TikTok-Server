package cache

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/pkg/tlog"
	"fmt"
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
}
