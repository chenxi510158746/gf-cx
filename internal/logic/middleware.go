package logic

import (
	"gf-cx/internal/model"
	"gf-cx/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func NewMiddleware() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// 初始化，务必最开始执行
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if userEntity := service.Session().GetUser(r.Context()); userEntity.Id > 0 {
		customCtx.User = &model.ContextUser{
			Id:       userEntity.Id,
			UserName: userEntity.UserName,
			Mobile:   userEntity.Mobile,
			Age:      userEntity.Age,
		}
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	user := service.Session().GetUser(r.Context())
	if user.Id > 0 {
		r.Middleware.Next()
	} else {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    http.StatusForbidden,
			Message: "未登陆",
		})
	}
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
