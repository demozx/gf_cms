package route

import (
	"gf_cms/internal/controller/backendApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台api路由分组
func backendApiHandle(s *ghttp.Server) {
	s.Group(service.Util().BackendApiGroup(), func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			service.Middleware().CORS,
		)
		group.ALLMap(g.Map{
			"/captcha/get": backendApi.Captcha.Get,
			"/admin/login": backendApi.Admin.Login,
		})
	})
	s.Group(service.Util().BackendApiGroup(), func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			service.Middleware().BackendAuthSession,
		)
		group.ALLMap(g.Map{
			"/admin/logout":      backendApi.Admin.Logout,
			"/admin/clear_cache": backendApi.Admin.ClearCache,
		})
	})
}
