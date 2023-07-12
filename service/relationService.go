package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
)

func RelationAction(authID, toUserID, actionType int64) (*message.DouyinRelationActionResponse, error) {
	var err error

	if actionType == 1 {
		err = model.FollowAction(authID, toUserID)
	}

	if actionType == 2 {
		err = model.UnFollowAction(authID, toUserID)
	}

	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	return &message.DouyinRelationActionResponse{}, nil
}

func GetFollowList(userID int64) (*message.DouyinRelationFollowListResponse, error) {

	followUserList, err := model.GetFollowList(userID)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	userList := PackUserList(followUserList)
	for i := range userList {
		userList[i].IsFollow = true
	}
	resp := &message.DouyinRelationFollowListResponse{
		UserList: userList,
	}

	return resp, nil
}

func GetFollowerList(userID int64) (*message.DouyinRelationFollowerListResponse, error) {

	followerList, err := model.GetFollowerList(userID)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	userList := PackUserList(followerList)
	for i := range userList {
		userList[i].IsFollow = true
	}
	resp := &message.DouyinRelationFollowerListResponse{
		UserList: userList,
	}

	return resp, nil
}

func PackUserList(users []*model.User) []*message.User {
	userList := make([]*message.User, len(users))
	for i, u := range users {
		userList[i] = &message.User{
			Id:              int64(u.ID),
			Name:            u.UserName,
			FollowCount:     u.FollowingCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        false,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		}
	}
	return userList
}
