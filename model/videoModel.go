package model

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	PublishTime   time.Time `gorm:"not null; index:idx_publish_time;"`
	Author        User      `gorm:"foreignKey:AuthorID;"`
	AuthorID      int64     `gorm:"index:idx_author_id; not null;"`
	PlayUrl       string    `gorm:"type:varchar(255); not null;"`
	CoverUrl      string    `gorm:"type:varchar(255); not null;"`
	FavoriteCount int64     `gorm:"default:0;"`
	CommentCount  int64     `gorm:"default:0;"`
	Title         string    `gorm:"type:varchar(64); not null;"`
}

func (Video) TableName() string {
	return "video"
}

func GetVideoList(videoIDs []int64) ([]*Video, error) {
	return nil, nil
}

func GetVideoListByUserID(userID int64) ([]*Video, error) {
	var videos []*Video
	author, err := GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	err = db.Model(&Video{}).Where(&Video{AuthorID: userID}).Find(&videos).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	for _, v := range videos {
		v.Author = *author
	}

	return videos, nil
}
