package cmd

import (
	"context"
	"gf-cx/internal/consts"
	"gf-cx/internal/controller"
	"gf-cx/internal/service"

	"github.com/gogf/gf/v2/net/goai"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				group.Bind(
					controller.Goods,
					controller.User.SignUp,
					controller.User.SignIn,
					controller.Game.Start,
					controller.Ocr.Do,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Order,
						controller.User.Info,
						controller.User.SignOut,
					)
				})
			})
			OpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func OpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`
	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Version:     consts.OpenAPIUserVersion,
	}
}
