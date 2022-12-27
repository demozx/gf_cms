package router

import (
	"gf_cms/internal/controller/pcApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func pcApiHandle(s *ghttp.Server) {
	s.Group(service.Util().PcApiGroup(), func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
		)
		// pcApi路由
		group.ALLMap(g.Map{
			"/guestbook.html": pcApi.Guestbook.Submit,
		})
	})
}
