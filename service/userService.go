package service

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/auth"
	"TikTokServer/pkg/errorcode"

	"golang.org/x/crypto/bcrypt"
)

func UserRegister(userName string, password string) (*message.DouyinUserRegisterResponse, error) {

	if len(userName) > 32 || len(password) > 32 {
		return nil, errorcode.ErrHttpReachMaxCount
	}
	if len(userName) == 0 || len(password) == 0 {
		return nil, errorcode.ErrHttpSecretEmptyData
	}
	user, err := model.QuaryUser(userName)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	if user != nil {
		return nil, errorcode.ErrHttpUserAlreadyExist
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		errCode := errorcode.ErrHttpEncrypt
		errCode.SetError(err)
		return nil, errCode
	}

	userInfo, err := model.CreateUser(userName, string(hashedPassword))
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	token, err := auth.CreateToken(int64(userInfo.ID), userName)
	if err != nil {
		return nil, err
	}

	resp := &message.DouyinUserRegisterResponse{
		UserId: int64(userInfo.ID),
		Token:  token,
	}
	return resp, nil
}

func UserLogin(userName string, password string) (*message.DouyinUserLoginResponse, error) {
	if len(userName) == 0 || len(password) == 0 {
		return nil, errorcode.ErrHttpSecretEmptyData
	}
	if len(userName) > 32 || len(password) > 32 {
		return nil, errorcode.ErrHttpReachMaxCount
	}
	user, err := model.QuaryUser(userName)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	if user == nil {
		return nil, errorcode.ErrHttpUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errorcode.ErrHttpPasswordIncorrect
	}

	token, err := auth.CreateToken(int64(user.ID), userName)
	if err != nil {
		return nil, err
	}

	resp := &message.DouyinUserLoginResponse{
		UserId: int64(user.ID),
		Token:  token,
	}

	return resp, nil
}

func GetUserInfo(userID int64) (*message.DouyinUserResponse, error) {

	user, err := model.GetUserByID(userID)

	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}

	resp := &message.DouyinUserResponse{
		User: &message.User{
			Id:              int64(user.ID),
			Name:            user.UserName,
			FollowCount:     int64(user.FollowingCount),
			FollowerCount:   int64(user.FollowerCount),
			IsFollow:        false,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  int64(len(user.FavoriteVideo)),
			WorkCount:       int64(user.WorkCount),
			FavoriteCount:   int64(len(user.FavoriteVideo)),
		},
	}
	return resp, nil
}
