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
func (c *cRouterBeautify) About(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.AboutChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// News 路由美化-新闻动态
func (c *cRouterBeautify) News(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// NewsCompany 路由美化-公司新闻
func (c *cRouterBeautify) NewsCompany(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsCompanyChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// NewsIndustry 路由美化-行业动态
func (c *cRouterBeautify) NewsIndustry(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsIndustryChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// Product 产品展示
func (c *cRouterBeautify) Product(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = Image.List(ctx, &pc.ImageListReq{ChannelId: consts.ProductChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// Honor 荣誉资质
func (c *cRouterBeautify) Honor(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = Image.List(ctx, &pc.ImageListReq{ChannelId: consts.HonorChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// Guestbook 在线留言
func (c *cRouterBeautify) Guestbook(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.GuestbookChannelTid})
	if err != nil {
		return nil, err
	}
	return
}

// Contact 联系我们
func (c *cRouterBeautify) Contact(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.RouterBeautifyRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.ContactChannelTid})
	if err != nil {
		return nil, err
	}
	return
}
