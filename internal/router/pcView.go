package router

import (
	"gf_cms/internal/controller/pc"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func pcViewHandle(s *ghttp.Server) {
	s.Group("", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().PcResponse,
		)
		// pc路由
		group.ALLMap(g.Map{
			"/":                         pc.Index.Index,       // 首页
			"/index.html":               pc.Index.Index,       // 首页
			"/article/list/{id}.html":   pc.Article.List,      // 文章列表
			"/article/detail/{id}.html": pc.Article.Detail,    // 文章详情
			"/single_page/{id}.html":    pc.SinglePage.Detail, // 单页
			"/image/list/{id}.html":     pc.Image.List,        // 图集列表

			"/news.html": pc.Article.List,
		})
	})
}
