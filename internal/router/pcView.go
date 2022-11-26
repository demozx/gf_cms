package router

import (
	"gf_cms/internal/controller/pc"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func pcViewHandle(s *ghttp.Server) {
	s.Group("", func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
		)
		group.ALLMap(g.Map{
			"/": pc.Index.Index,
		})
	})
}
