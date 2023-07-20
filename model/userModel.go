package model

import (
	"TikTokServer/pkg/tlog"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName        string  `gorm:"type:varchar(32);not null;index:idx_username,unique;"`
	Password        string  `gorm:"type:varchar(80);not null;"`
	FavoriteVideos  []Video `gorm:"many2many:user_favorite_videos;"`
	TotalFavorited  int64   `gorm:"default:0;"` // 投稿视频获赞总数
	WorkCount       int64   `gorm:"default:0;"`
	FollowingCount  int64   `gorm:"default:0;"`
	FollowerCount   int64   `gorm:"default:0;"`
	FavoriteCount   int64   `gorm:"default:0;"` // 点赞视频数
	Avatar          string  `gorm:"type:varchar(255);default:NULL;"`
	BackgroundImage string  `gorm:"type:varchar(255);default:NULL;"`
	Signature       string  `gorm:"type:varchar(255);default:NULL;"`
}

func (User) TableName() string {
	return "user"
}

func GetUserByID(userID int64) (*User, error) {
	user := &User{}
	result := db.Where("id = ?", userID).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	// tlog.Infof("User: %v", user)
	return user, nil
}

func MutilGetUserByID(userIDs []int64) ([]*User, error) {
	users := make([]*User, 0)
	if len(userIDs) == 0 {
		return users, nil
	}
	if err := db.Where("id in ?", userIDs).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(userName, password string) (*User, error) {
	user := &User{
		UserName:        userName,
		Password:        password,
		Avatar:          "https://oss.nikunokoya.com/blogbucket/images/others/avatar.png",
		BackgroundImage: "https://oss.nikunokoya.com/blogbucket/post/MMD-RayCast/RayEffect.png",
		Signature:       "这个人很懒，什么都没有留下",
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	tlog.Infof("CreateUser success: %v", user)
	return user, nil
}

func QuaryUserByName(userName string) (*User, error) {
	user := &User{}
	result := db.Where("user_name = ?", userName).Take(&user)
	// tlog.Debugf("QuaryUser: ", result.Error)
	// tlog.Debugf("user: %v", user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}
	return user, nil
}
