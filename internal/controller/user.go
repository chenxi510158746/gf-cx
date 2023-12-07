package controller

import (
	"context"
	v1 "gf-cx/api/v1"
	"gf-cx/internal/model"
	"gf-cx/internal/service"
)

var User = cUser{}

type cUser struct{}

func (c *cUser) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	userId, err := service.User().SignUp(ctx, model.UserSignUpInput{
		UserName:   req.UserName,
		Mobile:     req.Mobile,
		Password:   req.Password,
		RePassword: req.RePassword,
		Age:        req.Age,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserSignUpRes{UserId: userId}, err
}

func (c *cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	err = service.User().SignIn(ctx, model.UserSignInInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cUser) Info(ctx context.Context, req *v1.UserInfoReq) (res *v1.UserInfoRes, err error) {
	out, err := service.User().UserInfo(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoRes{User: out.User}, err
}

func (c *cUser) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	if err != nil {
		return nil, err
	}
	return
}
