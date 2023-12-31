package service

import (
	"context"
	"gf-cx/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IBizCtx interface {
	Init(r *ghttp.Request, customCtx *model.Context)
	Get(ctx context.Context) *model.Context
	SetUser(ctx context.Context, ctxUser *model.ContextUser)
}

var localBizCtx IBizCtx

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
