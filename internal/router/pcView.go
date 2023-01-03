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
		pcHost := service.Util().GetConfig("server.pcHost")
		if len(pcHost) > 0 {
			// 如果设置了pc域名，则路由需绑定域名访问
			pcHost = "@" + pcHost
		}
		// pc路由
		group.ALLMap(g.Map{
			"/" + pcHost:                         pc.Index.Index,       // 首页
			"/index.html" + pcHost:               pc.Index.Index,       // 首页
			"/article/list/{id}.html" + pcHost:   pc.Article.List,      // 文章列表
			"/article/detail/{id}.html" + pcHost: pc.Article.Detail,    // 文章详情
			"/image/list/{id}.html" + pcHost:     pc.Image.List,        // 图集列表
			"/image/detail/{id}.html" + pcHost:   pc.Image.Detail,      // 图集详情
			"/single_page/{id}.html" + pcHost:    pc.SinglePage.Detail, // 单页
			"/search.html" + pcHost:              pc.Search.Index,      // 搜索
			// 路由美化
			"/about.html" + pcHost:         pc.RouterBeautify.About,        // 公司简介
			"/news.html" + pcHost:          pc.RouterBeautify.News,         // 新闻动态
			"/news_company.html" + pcHost:  pc.RouterBeautify.NewsCompany,  // 公司新闻
			"/news_industry.html" + pcHost: pc.RouterBeautify.NewsIndustry, // 行业动态
			"/product.html" + pcHost:       pc.RouterBeautify.Product,      // 产品展示
			"/product_{id}.html" + pcHost:  pc.Image.List,                  // 子产品
			"/honor.html" + pcHost:         pc.RouterBeautify.Honor,        // 荣誉资质
			"/guestbook.html" + pcHost:     pc.RouterBeautify.Guestbook,    // 在线留言
			"/contact.html" + pcHost:       pc.RouterBeautify.Contact,      // 联系我们
		})
	})
}
