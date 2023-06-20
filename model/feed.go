package model

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	PublishTime   time.Time `gorm:"not null; index:idx_publish_time;"`
	Author        User      `gorm:"foreignKey:AuthorID;"`
	AuthorID      int       `gorm:"index:idx_author_id; not null;"`
	PlayUrl       string    `gorm:"type:varchar(255); not null;"`
	CoverUrl      string    `gorm:"type:varchar(255); not null;"`
	FavoriteCount int       `gorm:"default:0;"`
	CommentCount  int       `gorm:"default:0;"`
	Title         string    `gorm:"type:varchar(64); not null;"`
}

func (Video) TableName() string {
	return "video"
}

func GetVideoList(videoIDs []int64) ([]*Video, error) {
	return nil, nil
}
