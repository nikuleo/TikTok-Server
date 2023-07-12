package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Relation struct {
	gorm.Model
	User     User  `gorm:"foreignKey:UserID;"`
	UserID   int64 `gorm:"index:idx_userid,unique; not null;"`
	ToUser   User  `gorm:"foreignKey:UserID;"`
	ToUserID int64 `gorm:"index:idx_userid,unique; index:idx_to_userid; not null;"`
}

type Friend struct {
	gorm.Model
	User     User  `gorm:"foreignKey:UserID;"`
	UserID   int64 `gorm:"index:idx_userid,unique; not null;"`
	Friend   User  `gorm:"foreignKey:UserID;"`
	FriendID int64 `gorm:"index:idx_userid,unique; index:idx_friendid; not null;"`
}

func (Relation) TableName() string {
	return "relation"
}

func GetFollowList(uid int64) ([]*User, error) {

	var relations []*Relation
	err := db.Where("user_id = ?", uid).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	userList := make([]*User, len(relations))

	for i, relation := range relations {
		user, err := GetUserByID(relation.ToUserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		userList[i] = user
	}

	return userList, nil
}

func GetFollowerList(tid int64) ([]*User, error) {

	var relations []*Relation
	err := db.Where("to_user_id = ?", tid).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	userList := make([]*User, len(relations))

	for i, relation := range relations {
		user, err := GetUserByID(relation.UserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		userList[i] = user
	}

	return userList, nil
}

// uid关注tid，所以uid的关注人数加一，tid的粉丝数加一
func FollowAction(uid, tid int64) error {
	err := db.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(&Relation{UserID: uid, ToUserID: tid}).Error
		if err != nil {
			return err
		}

		//	判断是否互相关注
		isFriend := &Relation{}
		err = tx.Where("user_id = ? AND to_user_id = ?", tid, uid).First(isFriend).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		if isFriend.UserID == tid && isFriend.ToUserID == uid {
			// 互相关注，创建好友关系
			err = tx.Create(&Friend{UserID: uid, FriendID: tid}).Error
			if err != nil {
				return err
			}
			err = tx.Create(&Friend{UserID: tid, FriendID: uid}).Error
			if err != nil {
				return err
			}
		}
		res := tx.Model(&User{}).Where("id = ?", uid).Update("following_count", gorm.Expr("following_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return fmt.Errorf("数据库错误，查询记录不唯一")
		}

		res = tx.Model(&User{}).Where("id = ?", tid).Update("follower_count", gorm.Expr("follower_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return fmt.Errorf("数据库错误，查询记录不唯一")
		}
		return nil
	})
	return err
}

func UnFollowAction(uid, tid int64) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		relation := &Relation{}

		if err := tx.Where("user_id = ? AND to_user_id = ?", uid, tid).First(relation).Error; err != nil {
			return err
		}

		err := tx.Unscoped().Delete(relation).Error
		if err != nil {
			return err
		}

		// 更新 user 表中的 following count
		res := tx.Model(&User{}).Where("ID = ?", uid).Update("following_count", gorm.Expr("following_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return fmt.Errorf("数据库错误，查询记录不唯一")
		}

		// 更新 user 表中的 follower count
		res = tx.Model(&User{}).Where("ID = ?", tid).Update("follower_count", gorm.Expr("follower_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected != 1 {
			return fmt.Errorf("数据库错误，查询记录不唯一")
		}

		// 删除 friend 中的关系
		friend := &Friend{}
		if err := tx.Where("user_id = ? AND friend_id = ?", uid, tid).First(friend).Error; err != nil {
			return err
		}
		err = tx.Unscoped().Delete(friend).Error
		if err != nil {
			return err
		}
		mutuFriend := &Friend{}
		if err := tx.Where("user_id = ? AND friend_id = ?", tid, uid).First(mutuFriend).Error; err != nil {
			return err
		}
		err = tx.Unscoped().Delete(mutuFriend).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
