package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/consts"
)

var (
	RouterBeautify = cRouterBeautify{}
)

type cRouterBeautify struct{}

// About 路由美化-关于我们
func (c *cRouterBeautify) About(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.IndexRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.AbortChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// News 路由美化-新闻动态
func (c *cRouterBeautify) News(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.IndexRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// NewsCompany 路由美化-公司新闻
func (c *cRouterBeautify) NewsCompany(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.IndexRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: 2})
	if err != nil {
		return nil, err
	}
	return
}

// NewsIndustry 路由美化-行业动态
func (c *cRouterBeautify) NewsIndustry(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.IndexRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: 6})
	if err != nil {
		return nil, err
	}
	return
}
