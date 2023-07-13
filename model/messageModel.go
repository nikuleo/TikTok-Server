package model

import (
	"TikTokServer/pkg/tlog"
	"errors"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromUser   User   `gorm:"foreignKey:FromUserID;"` // 消息发送者
	FromUserID int64  `gorm:"index:idx_from_userid,unique; not null;"`
	ToUser     User   `gorm:"foreignKey:ToUserID;"` // 消息接收者
	ToUserID   int64  `gorm:"index:idx_from_userid,unique; index:idx_to_userid; not null;"`
	Content    string `gorm:"type:text; not null;"`
}

type FriendUser struct {
	UserInfo User
	Message  string `json:"message"` //聊天信息
	MsgType  int64  `json:"msgType"` //message信息的类型，0=>请求用户接受信息，1=>当前请求用户发送的信息
}

func (Message) TableName() string {
	return "message"
}

// 查询相互关注的强关注关系
func GetFriendList(uid int64) ([]FriendUser, error) {

	// var friends []*Friend
	friends := make([]*Friend, 0)
	err := db.Where("user_id = ?", uid).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	for _, friend := range friends {
		tlog.Debugf("Friend: friend: %+v", friend)
	}
	friendList := make([]FriendUser, len(friends))

	for i, u := range friends {
		tlog.Debugf("u.id : %v", u.ID)
		user, err := GetUserByID(u.FriendID)

		tlog.Debugf("Friend: user: %v, err: %v", user, err)

		if err != nil {
			return friendList, err
		}
		tlog.Debugf("userInfo: %+v", friendList[i].UserInfo)
		friendList[i].UserInfo = *user

		message, err := GetNewestMessageByUserID(u.FriendID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if err == gorm.ErrRecordNotFound {
			friendList[i].Message = ""
			friendList[i].MsgType = 0
			continue
		}
		friendList[i].Message = message.Content
		if message.FromUserID == uid {
			friendList[i].MsgType = 1
		} else {
			friendList[i].MsgType = 0
		}
	}

	return friendList, nil
}

func GetMessageListByID(fromUserID, toUserID int64) ([]*Message, error) {
	var messages []*Message
	resutl := db.Where("from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).Or("from_user_id = ? AND to_user_id = ?", toUserID, fromUserID).Order("created_at asc").Order("created_at asc").Find(&messages)
	if resutl.Error != nil {
		return nil, resutl.Error
	}

	return messages, nil
}

func CreateMessage(fromUserID, toUserID int64, content string) error {
	message := &Message{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Content:    content,
	}
	err := db.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}

func GetNewestMessageByUserID(userID int64) (*Message, error) {
	message := &Message{}

	result := db.Where("from_user_id = ? AND to_user_id = ?", userID, userID).Order("created_at desc").Limit(1).Find(&message)
	if result.Error != nil {
		return nil, result.Error
	}

	return message, nil
}

// func PackUserInfo(friendUser *FriendUser, user *User) {
// 	friendUser.UserName = user.UserName
// 	friendUser.TotalFavorited = user.TotalFavorited
// 	friendUser.WorkCount = user.WorkCount
// 	friendUser.FollowingCount = user.FollowingCount
// 	friendUser.FollowerCount = user.FollowerCount
// 	friendUser.FavoriteCount = user.FavoriteCount
// 	friendUser.Avatar = user.Avatar
// 	friendUser.BackgroundImage = user.BackgroundImage
// 	friendUser.Signature = user.Signature
// }
