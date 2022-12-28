package router

import (
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

func mobileApiHandle(s *ghttp.Server) {
	s.Group(service.Util().MApiGroup(), func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
		)

	})
}
