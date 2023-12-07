package service

import (
	"context"
	"gf-cx/internal/model"
)

type IUser interface {
	SignUp(ctx context.Context, req model.UserSignUpInput) (userId uint, err error)
	SignIn(ctx context.Context, req model.UserSignInInput) (err error)
	UserInfo(ctx context.Context) (out model.UserInfoOut, err error)
	SignOut(ctx context.Context) error
	IsUserNameAvailable(ctx context.Context, userName string) (b bool, err error)
	IsMobileAvailable(ctx context.Context, mobile string) (b bool, err error)
	EncryptPassword(userName string, password string) string
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
