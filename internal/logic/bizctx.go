package logic

import (
	"context"
	"gf-cx/internal/consts"
	"gf-cx/internal/model"
	"gf-cx/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sBizCtx struct{}

func init() {
	service.RegisterBizCtx(NewBizCtx())
}

func NewBizCtx() service.IBizCtx {
	return &sBizCtx{}
}

func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

func (s *sBizCtx) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func (s *sBizCtx) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}
