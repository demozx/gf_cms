package route

import (
	"gf_cms/internal/controller/adminApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台api路由分组
func backendApiHandle(s *ghttp.Server) {
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
			"/admin/logout":      adminApi.Admin.Logout,
			"/admin/clear_cache": adminApi.Admin.ClearCache,
		})
	})
}
