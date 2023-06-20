package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string  `gorm:"type:varchar(32); not null; index:idx_username,unique;"`
	Password        string  `gorm:"type:varchar(80); not null;"`
	FavoriteVideo   []Video `gorm:"many2many:user_favorite_videos;"`
	FollowingCount  int     `gorm:"default:0;"`
	FollowerCount   int     `gorm:"default:0;"`
	Avatar          string  `gorm:"type:varchar(255); default:NULL;"`
	BackgroundImage string  `gorm:"type:varchar(255); default:NULL;"`
	Signature       string  `gorm:"type:varchar(255); default:NULL;"`
}

func (User) TableName() string {
	return "user"
}

func GetUserByID(userID int64) (*User, error) {
	//TODO:
	return nil, nil
}

func CreateUser(user []*User) (err error) {
	return nil
}

func QuaryUser(userName string) ([]*User, error) {
	return nil, nil
}
