package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
)

type DouyinMessageListResponse struct {
	StatusCode  int32              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string             `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`       // 返回状态描述
	MessageList []*message.Message `protobuf:"bytes,3,opt,name=message_list,json=messageList,proto3" json:"message_list,omitempty"` // 消息列表
}

func GetFriendList(userID int64) (*message.DouyinRelationFriendListResponse, error) {
	friendList, err := model.GetFriendList(userID)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	resp := &message.DouyinRelationFriendListResponse{
		UserList: PackFriendList(friendList),
	}
	return resp, nil
}

func MessageAction(fromUserID, toUserID, actionType int64, content string) (*message.DouyinMessageActionResponse, error) {

	if actionType == 1 {
		err := model.CreateMessage(fromUserID, toUserID, content)
		if err != nil {
			errCode := errorcode.ErrHttpDatabase
			errCode.SetError(err)
			return nil, errCode
		}
	}

	return &message.DouyinMessageActionResponse{}, nil
}

func GetMessageList(fromUserID, toUserID int64) (*DouyinMessageListResponse, error) {
	messages, err := model.GetMessageListByID(fromUserID, toUserID)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	resp := &DouyinMessageListResponse{
		MessageList: PackMessageList(messages),
	}
	return resp, nil
}

func PackMessageList(messages []*model.Message) []*message.Message {
	messageList := make([]*message.Message, len(messages))
	for i, m := range messages {
		messageList[i] = &message.Message{
			Id:         int64(m.ID),
			ToUserId:   m.ToUserID,
			FromUserId: m.FromUserID,
			Content:    m.Content,
			CreateTime: m.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}
	return messageList
}

func PackFriendList(users []model.FriendUser) []*message.FriendUser {
	friendList := make([]*message.FriendUser, len(users))
	for i, u := range users {
		friendList[i] = &message.FriendUser{
			Message:         u.Message,
			MsgType:         u.MsgType,
			Id:              int64(u.UserInfo.ID),
			Name:            u.UserInfo.UserName,
			FollowCount:     u.UserInfo.FollowingCount,
			FollowerCount:   u.UserInfo.FollowerCount,
			IsFollow:        true,
			Avatar:          u.UserInfo.Avatar,
			BackgroundImage: u.UserInfo.BackgroundImage,
			Signature:       u.UserInfo.Signature,
			TotalFavorited:  u.UserInfo.TotalFavorited,
			WorkCount:       u.UserInfo.WorkCount,
			FavoriteCount:   u.UserInfo.FavoriteCount,
		}
	}
	return friendList
}
