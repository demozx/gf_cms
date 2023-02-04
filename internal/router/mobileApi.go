package router

import (
	"gf_cms/internal/controller/mobileApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func mobileApiHandle(s *ghttp.Server) {
	s.Group(service.Util().MApiGroup(), func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			service.Middleware().FilterXSS,
		)
		mobileHost := service.Util().GetConfig("server.mobileHost")
		if len(mobileHost) > 0 {
			// 如果设置了mobile域名，则路由需绑定域名访问
			mobileHost = "@" + mobileHost
		}
		// pcApi路由
		group.ALLMap(g.Map{
			"/guestbook.html" + mobileHost: mobileApi.Guestbook.Submit,
		})
	})
}
