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
			service.Middleware().FilterXSS,
			ghttp.MiddlewareGzip,
		)
		pcHost := service.Util().GetConfig("server.pcHost")
		if len(pcHost) > 0 {
			// 如果设置了pc域名，则路由需绑定域名访问
			pcHost = "@" + pcHost
		}
		// pcApi路由
		group.ALLMap(g.Map{
			"/guestbook.html" + pcHost: pcApi.Guestbook.Submit,
		})
	})
}
