package model

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	User       User  `gorm:"foreignKey:UserID;"`
	UserID     int64 `gorm:"index:idx_userid,unique; not null;"`
	ToUser     User  `gorm:"foreignKey:UserID;"`
	FollowerID int64 `gorm:"index:idx_userid,unique; index:idx_to_userid; not null;"`
}

func (Relation) TableName() string {
	return "relation"
}

func GetFriendList(uid, tid int64) (*Relation, error) {
	// TODO: relation 中相互关注的用户
	return nil, nil
}

func GetFollowingList(uid int64) ([]*Relation, error) {
	return nil, nil
}

func GetFollowerList(tid int64) ([]*Relation, error) {
	return nil, nil
}

func FollowAction(uid, tid int64) error {
	return nil
}

func UnFollowAction(uid, tid int64) error {
	return nil
}
