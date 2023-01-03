package router

import (
	"gf_cms/internal/controller/mobile"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func mobileViewHandle(s *ghttp.Server) {
	s.Group("", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().MobileResponse,
		)
		mobileHost := service.Util().GetConfig("server.mobileHost")
		if len(mobileHost) > 0 {
			// 如果设置了mobile域名，则路由需绑定域名访问
			mobileHost = "@" + mobileHost
			// mobile路由
			group.ALLMap(g.Map{
				"/" + mobileHost:                      mobile.Index.Index,       // 首页
				"/index.html" + mobileHost:            mobile.Index.Index,       // 首页
				"/single_page/{id}.html" + mobileHost: mobile.SinglePage.Detail, // 单页
				// 路由美化
				"/about.html" + mobileHost: mobile.RouterBeautify.About, // 公司简介
			})
		}
	})
}
