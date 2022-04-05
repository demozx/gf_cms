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
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"
	"time"
)

var Ctx = gctx.New()

// 项目ProjectName
var ProjectName, _ = g.Cfg().Get(Ctx, "server.projectName", "gf_cms")

//后台入口前缀
var AdminPrefix, _ = g.Cfg().Get(Ctx, "server.adminPrefix", "admin")

//AdminGroup 后台view分组
var AdminGroup = "/" + AdminPrefix.String()

//AdminApiGroup 后台api分组
var AdminApiGroup = "/" + AdminPrefix.String() + "_api" //system_config表根据name获取值
func funSystemConfig(name string) string {
	cacheKey := ProjectName.String() + ":system_config:" + name
	exists, err := g.Redis().Do(Ctx, "EXISTS", cacheKey)
	if err != nil {
		panic(err)
	}
	//存在缓存key
	if exists.Bool() {
		value, err := g.Redis().Do(Ctx, "GET", cacheKey)
		if err != nil {
			panic(err)
		}
		return value.String()
	}
	//不存在缓存key，从数据库读取
	val, _ := g.Model("system_config").Where("name", name).Value("value")
	g.Redis().Do(Ctx, "SET", cacheKey, val.String())
	return val.String()
}

//生成后台view的url
func funAdminUrl(route string) string {
	return AdminGroup + route
}

//生成后台api的url
func funAdminApiUrl(route string) string {
	return AdminApiGroup + route
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 给模板视图全局绑定方法
			g.View().BindFuncMap(g.Map{
				"system_config": funSystemConfig,
				"admin_url":     funAdminUrl,
				"admin_api_url": funAdminApiUrl,
			})
			s := g.Server()
			//session使用redis
			s.SetConfigWithMap(g.Map{
				"SessionMaxAge":  time.Minute,
				"SessionStorage": gsession.NewStorageRedis(g.Redis(), ProjectName.String()+":"+consts.AdminSessionKeyPrefix+":"),
			})
			/*后台view路由分组*/
			s.Group(AdminGroup, func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
				)
				group.ALLMap(g.Map{
					"/admin/login": admin.Admin.Login,
				})
			})
			s.Group(AdminGroup, func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().AuthSession,
				)
				group.ALLMap(g.Map{
					"/": admin.Index.Index,
				})
			})
			/*后台api路由分组*/
			s.Group(AdminApiGroup, func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().CORS,
				)
				group.ALLMap(g.Map{
					"/captcha/get": adminApi.Captcha.Get,
					"/admin/login": adminApi.Admin.Login,
				})
			})
			s.Group(AdminApiGroup, func(group *ghttp.RouterGroup) {
				group.Middleware(
					ghttp.MiddlewareHandlerResponse,
					service.Middleware().AuthSession,
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
