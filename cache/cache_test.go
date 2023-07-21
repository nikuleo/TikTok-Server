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
		Id:              0,
		Name:            "niku",
		FollowCount:     0,
		FollowerCount:   0,
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

func TestCommentCache(t *testing.T) {
	comments := []*message.Comment{}
	for i := 0; i < 10; i++ {
		comment := &message.Comment{
			Id:         int64(i),
			User:       nil,
			Content:    "comment" + strconv.Itoa(i),
			CreateDate: "2021-01-01",
		}
		comments = append(comments, comment)
	}

	err := SetVideoCommentToCache(1, comments)
	if err != nil {
		t.Error(err)
	}

	getComments, err := GetVideoCommentFromCache(1)
	if err != nil {
		t.Error(err)
	}
	for _, comment := range getComments {
		fmt.Printf("comment: %+v \n", comment)
	}
	// fmt.Printf("getComments: %+v \n", getComments)
}
