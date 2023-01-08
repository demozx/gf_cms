package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/consts"
)

var (
	RouterBeautify = cRouterBeautify{}
)

type cRouterBeautify struct{}

// About 路由美化-关于我们
func (c *cRouterBeautify) About(ctx context.Context, req *mobile.RouterBeautifyDefaultReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = SinglePage.Detail(ctx, &mobile.SinglePageReq{Id: consts.AboutChannelId})
	if err != nil {
		return nil, err
	}
	return
}

// Contact 路由美化-联系我们
func (c *cRouterBeautify) Contact(ctx context.Context, req *mobile.RouterBeautifyDefaultReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = SinglePage.Detail(ctx, &mobile.SinglePageReq{Id: consts.ContactChannelId})
	if err != nil {
		return nil, err
	}
	return
}

// Guestbook 路由美化-在线留言
func (c *cRouterBeautify) Guestbook(ctx context.Context, req *mobile.RouterBeautifyDefaultReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = SinglePage.Detail(ctx, &mobile.SinglePageReq{Id: consts.GuestbookChannelId})
	if err != nil {
		return nil, err
	}
	return
}

// News 路由美化-新闻动态
func (c *cRouterBeautify) News(ctx context.Context, req *mobile.RouterBeautifyPageSizeReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Article.List(ctx, &mobile.ArticleListReq{ChannelId: consts.NewsChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// NewsCompany 路由美化-公司新闻
func (c *cRouterBeautify) NewsCompany(ctx context.Context, req *mobile.RouterBeautifyPageSizeReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Article.List(ctx, &mobile.ArticleListReq{ChannelId: consts.NewsCompanyChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// NewsIndustry 路由美化-行业动态
func (c *cRouterBeautify) NewsIndustry(ctx context.Context, req *mobile.RouterBeautifyPageSizeReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Article.List(ctx, &mobile.ArticleListReq{ChannelId: consts.NewsIndustryChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// NewsDetail 路由美化-新闻详情
func (c *cRouterBeautify) NewsDetail(ctx context.Context, req *mobile.RouterBeautifyDetailReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Article.Detail(ctx, &mobile.ArticleDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// Product 路由美化-产品展示
func (c *cRouterBeautify) Product(ctx context.Context, req *mobile.RouterBeautifyPageSizeReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Image.List(ctx, &mobile.ImageListReq{ChannelId: consts.ProductChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// ProductDetail 路由美化-产品详情
func (c *cRouterBeautify) ProductDetail(ctx context.Context, req *mobile.RouterBeautifyDetailReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Image.Detail(ctx, &mobile.ImageDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// Honor 路由美化-荣誉资质
func (c *cRouterBeautify) Honor(ctx context.Context, req *mobile.RouterBeautifyDefaultReq) (res *mobile.RouterBeautifyDefaultRes, err error) {
	_, err = Image.List(ctx, &mobile.ImageListReq{ChannelId: consts.HonorChannelId})
	if err != nil {
		return nil, err
	}
	return
}
