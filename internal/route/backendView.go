package route

import (
	"gf_cms/internal/controller/backend"
	"gf_cms/internal/logic/middleware"
	"gf_cms/internal/logic/util"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台view路由分组
func backendViewHandle(s *ghttp.Server) {
	var backendGroup = util.Util().BackendGroup()
	s.Group(backendGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
		)
		group.ALLMap(g.Map{
			"/admin/login": backend.Admin.Login,
		})
	})
	s.Group(backendGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			middleware.Middleware().BackendAuthSession,
			middleware.Middleware().BackendCheckPolicy,
		)
		group.ALLMap(g.Map{
			"/":             backend.Index.Index,   //后台首页
			"welcome/index": backend.Welcome.Index, //后台欢迎页
			"channel/index": backend.Channel.Index, //栏目分类
			"setting/index": backend.Setting.Index, //后台设置
			"admin/index":   backend.Admin.Index,   // 管理员列表
		})
	})
}
