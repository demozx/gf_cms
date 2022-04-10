package cmd

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/controller/admin"
	"gf_cms/internal/controller/adminApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 给模板视图全局绑定方法
			g.View().BindFuncMap(g.Map{
				"system_config": service.ViewBindFun().SystemConfig,
				"admin_url":     service.ViewBindFun().AdminUrl,
				"admin_api_url": service.ViewBindFun().AdminApiUrl,
			})
			s := g.Server()
			//session使用redis
			s.SetConfigWithMap(g.Map{
				// session一天过期
				"SessionMaxAge":  time.Hour * 24,
				"SessionStorage": gsession.NewStorageRedis(g.Redis(), service.Util().ProjectName()+":"+consts.AdminSessionKeyPrefix+":"),
			})
			//路由
			/*后台view路由分组*/
			s.Group(service.Util().AdminGroup(), func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
				)
				group.ALLMap(g.Map{
					"/admin/login": admin.Admin.Login,
				})
			})
			s.Group(service.Util().AdminGroup(), func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().AdminAuthSession,
				)
				group.ALLMap(g.Map{
					"/": admin.Index.Index,
				})
			})
			/*后台api路由分组*/
			s.Group(service.Util().AdminApiGroup(), func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().CORS,
				)
				group.ALLMap(g.Map{
					"/captcha/get": adminApi.Captcha.Get,
					"/admin/login": adminApi.Admin.Login,
				})
			})
			s.Group(service.Util().AdminApiGroup(), func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().AdminAuthSession,
				)
				group.ALLMap(g.Map{
					"/admin/logout": adminApi.Admin.Logout,
				})
			})
			s.Run()
			return nil
		},
	}
)
