package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Video   Video  `gorm:"foreignKey:VideoID;"`
	VideoID int    `gorm:"index:idx_videoid; not null;"`
	User    User   `gorm:"foreignKey:UserID;"`
	UserID  int    `gorm:"index:idx_userid;not null;"`
	Content string `gorm:"type:varchar(255); not null"`
}

func (Comment) TableName() string {
	return "comment"
}

func NewComment(comment *Comment) (err error) {
	return nil
}

func GetCommentList(videoID int64) ([]*Comment, error) {
	return nil, nil
}

func DeleteComment(commentID, videoID int64) error {
	return nil
}
