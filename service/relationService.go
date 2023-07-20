package service

import (
	"TikTokServer/cache"
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
	// TODO: 使用消息队列异步删除缓存
	return &message.DouyinRelationActionResponse{}, nil
}

// 缓存优化循环查询
func GetFollowList(userID int64) (*message.DouyinRelationFollowListResponse, error) {

	// 从缓存中获取关注列表
	userList, err := getFollowListFromCache(userID)
	if err != nil {
		errCode := errorcode.ErrHttpCache
		errCode.SetError(err)
		return nil, errCode
	}

	// 缓存未命中，从数据库中获取关注列表，并写入缓存
	if userList == nil {
		userList, err = getFollowListFromDatabase(userID)
		if err != nil {
			errCode := errorcode.ErrHttpDatabase
			errCode.SetError(err)
			return nil, errCode
		}
	}

	cache.SetUserFollowing(userID, userList)

	resp := &message.DouyinRelationFollowListResponse{
		UserList: userList,
	}

	return resp, nil
}

// 通过缓存的 ID 优化循环查询
func GetFollowListByUserIDs(userIDs []int64) ([]*message.User, error) {

	followUserList, err := model.MutilGetUserByID(userIDs)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	userList := PackUserList(followUserList)
	for i := range userList {
		userList[i].IsFollow = true
	}
	return userList, nil
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

func getFollowListFromCache(userID int64) ([]*message.User, error) {
	followIDList, err := cache.GetUserFollowing(userID)
	if err != nil {
		return nil, err
	}
	if len(followIDList) == 0 {
		return nil, nil
	}

	return GetFollowListByUserIDs(followIDList)
}

func getFollowListFromDatabase(userID int64) ([]*message.User, error) {
	followUserList, err := model.GetFollowList(userID)
	if err != nil {
		return nil, err
	}

	userList := PackUserList(followUserList)
	for i := range userList {
		userList[i].IsFollow = true
	}

	return userList, nil
}

func getFollowUserIDs(userID int64) ([]int64, error) {
	followIDList, err := cache.GetUserFollowing(userID)
	if err != nil {
		return nil, err
	}
	if len(followIDList) != 0 {

		return followIDList, nil
	}
	userList, err := getFollowListFromDatabase(userID)
	if err != nil {
		return nil, err
	}
	followIDList = make([]int64, len(userList))
	for i, u := range userList {
		followIDList[i] = u.Id
	}
	return followIDList, nil
}
