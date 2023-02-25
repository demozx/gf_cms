package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/consts"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) Index(ctx context.Context, req *mobile.IndexReq) (res *mobile.IndexRes, err error) {
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, 0)
		chNavigation <- navigation
	}()
	// 公司简介栏目url
	chAboutChannelUrl := make(chan string, 1)
	go func() {
		defer close(chAboutChannelUrl)
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.AboutChannelId, "")
		chAboutChannelUrl <- url
	}()
	// 新闻动态栏目url
	chNewsChannelUrl := make(chan string, 1)
	go func() {
		defer close(chNewsChannelUrl)
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.NewsChannelId, "")
		chNewsChannelUrl <- url
	}()
	// 产品展示栏目url
	chProductChannelUrl := make(chan string, 1)
	go func() {
		defer close(chProductChannelUrl)
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ProductChannelId, "")
		chProductChannelUrl <- url
	}()
	// 联系我们栏目url
	chContactChannelUrl := make(chan string, 1)
	go func() {
		defer close(chContactChannelUrl)
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ContactChannelId, "")
		chContactChannelUrl <- url
	}()
	// 在线留言栏目url
	chGuestbookChannelUrl := make(chan string, 1)
	go func() {
		defer close(chGuestbookChannelUrl)
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookChannelUrl <- url
	}()
	// 荣誉资质栏目url
	chHonorChannelUrl := make(chan string, 1)
	go func() {
		defer close(chHonorChannelUrl)
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.HonorChannelId, "")
		chHonorChannelUrl <- url
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chGoodsChannelList)
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
	}()
	// banner广告
	chAdList := make(chan []*entity.CmsAd, 1)
	go func() {
		defer close(chAdList)
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.MobileHomeAdChannelId)
		chAdList <- adList
	}()
	// 产品展示
	chRecommendGoodsList := make(chan []*model.ImageListItem, 1)
	go func() {
		defer close(chRecommendGoodsList)
		goodsList, _ := service.Image().MobileHomeRecommendGoodsList(ctx, consts.ProductChannelId)
		chRecommendGoodsList <- goodsList
	}()
	// 关于我们简介
	chAboutChannelDescription := make(chan string, 1)
	go func() {
		defer close(chAboutChannelDescription)
		channel, _ := service.Channel().HomeAboutChannel(ctx, consts.AboutChannelId)
		chAboutChannelDescription <- channel.Description
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chTextNewsList)
		textNewsList, _ := service.Article().MobileHomeTextNewsList(ctx, consts.NewsChannelId)
		chTextNewsList <- textNewsList
	}()
	// 最新资讯-图片新闻
	chPicNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chPicNewsList)
		picNewsList, _ := service.Article().MobileHomePicNewsList(ctx, consts.NewsChannelId)
		chPicNewsList <- picNewsList
	}()
	// 荣誉资质列表
	chHonorList := make(chan []*model.ImageListItem, 1)
	go func() {
		defer close(chHonorList)
		honerList, _ := service.Image().MobileHomeHonerList(ctx, consts.HonorChannelId)
		chHonorList <- honerList
	}()

	err = service.Response().View(ctx, "/mobile/index/index.html", g.Map{
		"navigation":              <-chNavigation,              // 导航
		"aboutChannelUrl":         <-chAboutChannelUrl,         // 关于我们栏目url
		"newsChannelUrl":          <-chNewsChannelUrl,          // 新闻动态栏目url
		"productChannelUrl":       <-chProductChannelUrl,       // 产品展示栏目url
		"contactChannelUrl":       <-chContactChannelUrl,       // 联系我们栏目url
		"guestbookChannelUrl":     <-chGuestbookChannelUrl,     // 在线留言栏目url
		"honorChannelUrl":         <-chHonorChannelUrl,         // 荣誉资质栏目url
		"goodsChannelList":        <-chGoodsChannelList,        // 产品中心列表
		"adList":                  <-chAdList,                  // banner广告
		"recommendGoodsList":      <-chRecommendGoodsList,      // 产品列表
		"aboutChannelDescription": <-chAboutChannelDescription, // 关于我们简介
		"textNewsList":            <-chTextNewsList,            // 文字新闻
		"picNewsList":             <-chPicNewsList,             // 图片新闻
		"honorList":               <-chHonorList,               // 荣誉资质列表
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}
