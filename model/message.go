package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromUser   User   `gorm:"foreignKey:FromUserID;"` // 消息发送者
	FromUserID int    `gorm:"index:idx_from_userid,unique; not null;"`
	ToUser     User   `gorm:"foreignKey:ToUserID;"` // 消息接收者
	ToUserID   int    `gorm:"index:idx_from_userid,unique; index:idx_to_userid; not null;"`
	Content    string `gorm:"type:text; not null;"`
}

func (Message) TableName() string {
	return "message"
}

func GetMessageListByUserID(userID int64) ([]*Message, error) {
	return nil, nil
}

func SendMessage(fromUserID, toUserID int64, content string) error {
	return nil
}
