package model

import (
	"TikTokServer/pkg/errorcode"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Video   Video  `gorm:"foreignKey:VideoID;"`
	VideoID int64  `gorm:"index:idx_videoid; not null;"`
	User    User   `gorm:"foreignKey:UserID;"`
	UserID  int64  `gorm:"index:idx_userid;not null;"`
	Content string `gorm:"type:varchar(255); not null"`
}

func (Comment) TableName() string {
	return "comment"
}

func CreateComment(userID, videoID int64, commentText string) error {

	err := db.Transaction(func(tx *gorm.DB) error {
		comment := &Comment{
			VideoID: videoID,
			UserID:  userID,
			Content: commentText,
		}

		err := tx.Create(comment).Error
		if err != nil {
			return err
		}
		res := tx.Model(&Video{}).Where("ID = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))

		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errorcode.ErrHttpDatabase
		}
		return nil
	})

	return err
}

func DeleteComment(commentID, videoID int64) error {
	err := db.Transaction(func(tx *gorm.DB) error {

		comment := &Comment{}

		if err := tx.First(comment, commentID).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Delete(comment).Error; err != nil {
			return err
		}

		res := tx.Unscoped().Model(&Video{}).Where("ID = ?", videoID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return errorcode.ErrHttpDatabase
		}

		return nil
	})
	return err
}

func GetCommentList(videoID int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Model(&Comment{}).Where(&Comment{VideoID: videoID}).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	for i, c := range comments {
		user, err := GetUserByID(c.UserID)
		if err != nil {
			return comments, err
		}
		comments[i].User = *user
	}

	return comments, nil
}
