package route

import (
	"gf_cms/internal/controller/admin"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台view路由分组
func backendViewHandle(s *ghttp.Server) {
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
			service.Middleware().AdminCheckPolicy,
		)
		group.ALLMap(g.Map{
			"/": admin.Index.Index,
		})
	})
}
