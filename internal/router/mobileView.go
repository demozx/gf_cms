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
				"/" + mobileHost:                         mobile.Index.Index,       // 首页
				"/index.html" + mobileHost:               mobile.Index.Index,       // 首页
				"/article/list/{id}.html" + mobileHost:   mobile.Article.List,      // 文章列表
				"/article/detail/{id}.html" + mobileHost: mobile.Article.Detail,    // 文章详情
				"/image/list/{id}.html" + mobileHost:     mobile.Image.List,        // 图集列表
				"/image/detail/{id}.html" + mobileHost:   mobile.Image.Detail,      // 图集详情
				"/single_page/{id}.html" + mobileHost:    mobile.SinglePage.Detail, // 单页
				"/search.html" + mobileHost:              mobile.Search.Index,      // 搜索
				// 路由美化
				"/about/" + mobileHost:                   mobile.RouterBeautify.About,              // 公司简介
				"/about/index.html" + mobileHost:         mobile.RouterBeautify.About,              // 公司简介
				"/contact/" + mobileHost:                 mobile.RouterBeautify.Contact,            // 联系我们
				"/contact/index.html" + mobileHost:       mobile.RouterBeautify.Contact,            // 联系我们
				"/guestbook/" + mobileHost:               mobile.RouterBeautify.Guestbook,          // 在线留言
				"/guestbook/index.html" + mobileHost:     mobile.RouterBeautify.Guestbook,          // 在线留言
				"/news/" + mobileHost:                    mobile.RouterBeautify.News,               // 新闻动态
				"/news/index.html" + mobileHost:          mobile.RouterBeautify.News,               // 新闻动态
				"/news_company/" + mobileHost:            mobile.RouterBeautify.NewsCompany,        // 公司新闻
				"/news_company/index.html" + mobileHost:  mobile.RouterBeautify.NewsCompany,        // 公司新闻
				"/news_industry/" + mobileHost:           mobile.RouterBeautify.NewsIndustry,       // 行业动态
				"/news_industry/index.html" + mobileHost: mobile.RouterBeautify.NewsIndustry,       // 行业动态
				"/news/{id}.html" + mobileHost:           mobile.RouterBeautify.NewsDetail,         // 新闻详情
				"/news_company/{id}.html" + mobileHost:   mobile.RouterBeautify.NewsCompanyDetail,  // 公司新闻详情
				"/news_industry/{id}.html" + mobileHost:  mobile.RouterBeautify.NewsIndustryDetail, // 行业动态详情
				"/product/" + mobileHost:                 mobile.RouterBeautify.Product,            // 产品展示
				"/product/index.html" + mobileHost:       mobile.RouterBeautify.Product,            // 产品展示
				"/product_{id}.html" + mobileHost:        mobile.Image.List,                        // 子产品
				"/product/{id}.html" + mobileHost:        mobile.RouterBeautify.ProductDetail,      // 产品详情
				"/honor/" + mobileHost:                   mobile.RouterBeautify.Honor,              // 荣誉资质
				"/honor/index.html" + mobileHost:         mobile.RouterBeautify.Honor,              // 荣誉资质
				"/honor/{id}.html" + mobileHost:          mobile.RouterBeautify.HonorDetail,        // 荣誉资质详情
			})
		}
	})
}
