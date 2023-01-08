package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	RouterBeautify = cRouterBeautify{}
)

type cRouterBeautify struct{}

// About 路由美化-关于我们
func (c *cRouterBeautify) About(ctx context.Context, req *pc.RouterBeautifyDefaultReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.AboutChannelId})
	if err != nil {
		return nil, err
	}
	return
}

// News 路由美化-新闻动态
func (c *cRouterBeautify) News(ctx context.Context, req *pc.RouterBeautifyPageSizeReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	g.Dump(req)
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// NewsCompany 路由美化-公司新闻
func (c *cRouterBeautify) NewsCompany(ctx context.Context, req *pc.RouterBeautifyPageSizeReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsCompanyChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// NewsIndustry 路由美化-行业动态
func (c *cRouterBeautify) NewsIndustry(ctx context.Context, req *pc.RouterBeautifyPageSizeReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Article.List(ctx, &pc.ArticleListReq{ChannelId: consts.NewsIndustryChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// NewsDetail 路由美化-新闻详情
func (c *cRouterBeautify) NewsDetail(ctx context.Context, req *pc.RouterBeautifyDetailReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Article.Detail(ctx, &pc.ArticleDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// NewsCompanyDetail 路由美化-公司新闻详情
func (c *cRouterBeautify) NewsCompanyDetail(ctx context.Context, req *pc.RouterBeautifyDetailReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Article.Detail(ctx, &pc.ArticleDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// NewsIndustryDetail 路由美化-行业动态详情
func (c *cRouterBeautify) NewsIndustryDetail(ctx context.Context, req *pc.RouterBeautifyDetailReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Article.Detail(ctx, &pc.ArticleDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// Product 路由美化-产品展示
func (c *cRouterBeautify) Product(ctx context.Context, req *pc.RouterBeautifyPageSizeReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Image.List(ctx, &pc.ImageListReq{ChannelId: consts.ProductChannelId, Page: req.Page, Size: req.Size})
	if err != nil {
		return nil, err
	}
	return
}

// ProductDetail 路由美化-产品详情
func (c *cRouterBeautify) ProductDetail(ctx context.Context, req *pc.RouterBeautifyDetailReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Image.Detail(ctx, &pc.ImageDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// Honor 路由美化-荣誉资质
func (c *cRouterBeautify) Honor(ctx context.Context, req *pc.RouterBeautifyDefaultReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Image.List(ctx, &pc.ImageListReq{ChannelId: consts.HonorChannelId})
	if err != nil {
		return nil, err
	}
	return
}

// HonorDetail 路由美化-荣誉资质详情
func (c *cRouterBeautify) HonorDetail(ctx context.Context, req *pc.RouterBeautifyDetailReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = Image.Detail(ctx, &pc.ImageDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}

// Guestbook 路由美化-在线留言
func (c *cRouterBeautify) Guestbook(ctx context.Context, req *pc.RouterBeautifyDefaultReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.GuestbookChannelId})
	if err != nil {
		return nil, err
	}
	return
}

// Contact 路由美化-联系我们
func (c *cRouterBeautify) Contact(ctx context.Context, req *pc.RouterBeautifyDefaultReq) (res *pc.RouterBeautifyDefaultRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.ContactChannelId})
	if err != nil {
		return nil, err
	}
	return
}
