package route

import (
	"gf_cms/internal/controller/backendApi"
	"gf_cms/internal/logic/middleware"
	"gf_cms/internal/logic/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台api路由分组
func backendApiHandle(s *ghttp.Server) {
	var backendApiGroup = util.Util().BackendApiGroup()
	s.Group(backendApiGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			middleware.Middleware().CORS,
		)
		group.ALLMap(g.Map{
			"/captcha/get": backendApi.Captcha.Get,
			"/admin/login": backendApi.Admin.Login,
		})
	})
	s.Group(backendApiGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			middleware.Middleware().BackendAuthSession,
		)
		group.ALLMap(g.Map{
			"/admin/logout":      backendApi.Admin.Logout,
			"/admin/clear_cache": backendApi.Admin.ClearCache,
			"/welcome/index":     backendApi.Welcome.Index,
			"/setting/save":      backendApi.Setting.Save,
		})
	})
}
