package model

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/util"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	PublishTime   int64  `gorm:"not null; index:idx_publish_time;"`
	Author        User   `gorm:"foreignKey:AuthorID;"`
	AuthorID      int64  `gorm:"index:idx_author_id; not null;"`
	PlayUrl       string `gorm:"type:varchar(255); not null;"`
	CoverUrl      string `gorm:"type:varchar(255); not null;"`
	FavoriteCount int64  `gorm:"default:0;"`
	CommentCount  int64  `gorm:"default:0;"`
	Title         string `gorm:"type:varchar(64); not null;"`
}

func (Video) TableName() string {
	return "video"
}

func GetVideoListByTime(latestTime int64, limit int) ([]*Video, error) {
	videos := make([]*Video, 0)

	if err := db.Limit(limit).Order("publish_time desc").Find(&videos, "publish_time < ?", latestTime).Error; err != nil {
		return nil, err
	}

	for i, v := range videos {
		author, err := GetUserByID(v.AuthorID)
		if err != nil {
			return videos, err
		}
		videos[i].Author = *author
	}

	return videos, nil
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

func CreateVideo(userID int64, videoUrl, coverUrl, title string) error {
	video := &Video{
		AuthorID:      userID,
		PlayUrl:       videoUrl,
		CoverUrl:      coverUrl,
		Title:         title,
		PublishTime:   util.GetCurrentTime(),
		FavoriteCount: 0,
		CommentCount:  0,
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(video).Error
		if err != nil {
			return err
		}

		res := tx.Model(&User{}).Where("ID = ?", userID).Update("work_count", gorm.Expr("work_count + ?", 1))

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			errCode := errorcode.ErrHttpDatabase
			errCode.Msg = "插入视频时查询到的用户数不为1"
			return errCode
		}

		return nil
	})
	return err
}
