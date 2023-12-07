package service

import (
	"context"
	"gf-cx/internal/model/entity"
)

type ISession interface {
	SetUser(ctx context.Context, user *entity.Users) error
	GetUser(ctx context.Context) *entity.Users
	RemoveUser(ctx context.Context) error
}

var localSession ISession

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
