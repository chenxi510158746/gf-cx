package logic

import (
	"context"
	"gf-cx/internal/consts"
	"gf-cx/internal/model/entity"
	"gf-cx/internal/service"
)

type sSession struct{}

func init() {
	service.RegisterSession(NewSession())
}

func NewSession() service.ISession {
	return &sSession{}
}

func (s *sSession) SetUser(ctx context.Context, user *entity.Users) error {
	return service.BizCtx().Get(ctx).Session.Set(consts.SessionKeyUser, user)
}

func (s *sSession) GetUser(ctx context.Context) *entity.Users {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(consts.SessionKeyUser)
		if !v.IsNil() {
			var user *entity.Users
			_ = v.Struct(&user)
			return user
		}
	}
	return &entity.Users{}
}

func (s *sSession) RemoveUser(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.SessionKeyUser)
	}
	return nil
}
