package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/consts"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) Index(ctx context.Context, req *mobile.IndexReq) (res *mobile.IndexRes, err error) {
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().Navigation(ctx, 0)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "mobile导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNavigation <- navigation
		defer close(chNavigation)
	}()
	// 公司简介栏目url
	chAboutChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.AboutChannelId, "")
		chAboutChannelUrl <- url
		defer close(chAboutChannelUrl)
	}()
	// 新闻动态栏目url
	chNewsChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.NewsChannelId, "")
		chNewsChannelUrl <- url
		defer close(chNewsChannelUrl)
	}()
	// 产品展示栏目url
	chProductChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ProductChannelId, "")
		chProductChannelUrl <- url
		defer close(chProductChannelUrl)
	}()
	// 联系我们栏目url
	chContactChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ContactChannelId, "")
		chContactChannelUrl <- url
		defer close(chContactChannelUrl)
	}()
	// 在线留言栏目url
	chGuestbookChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookChannelUrl <- url
		defer close(chGuestbookChannelUrl)
	}()
	// 荣誉资质栏目url
	chHonorChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.HonorChannelId, "")
		chHonorChannelUrl <- url
		defer close(chHonorChannelUrl)
	}()
	// 附件栏目url
	chProductFuJianChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ProductFujianChannelId, "")
		chProductFuJianChannelUrl <- url
		defer close(chProductFuJianChannelUrl)
	}()
	// 钢背附件栏目url
	chProductGangBeiChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ProductGangBeiChannelId, "")
		chProductGangBeiChannelUrl <- url
		defer close(chProductGangBeiChannelUrl)
	}()
	// 减震片栏目url
	chProductJianZhenPianChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().ChannelUrl(ctx, consts.ProductJianZhenPianChannelId, "")
		chProductJianZhenPianChannelUrl <- url
		defer close(chProductJianZhenPianChannelUrl)
	}()
	// banner广告
	chAdList := make(chan []*entity.CmsAd, 1)
	go func() {
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.MobileHomeAdChannelId)
		chAdList <- adList
		defer close(chAdList)
	}()
	// 产品展示
	chRecommendGoodsList := make(chan []*model.ImageListItem, 1)
	go func() {
		goodsList, _ := service.Image().MobileHomeRecommendGoodsList(ctx, consts.ProductChannelId)
		chRecommendGoodsList <- goodsList
		defer close(chRecommendGoodsList)
	}()
	// 关于我们简介
	chAboutChannelDescription := make(chan string, 1)
	go func() {
		channel, _ := service.Channel().HomeAboutChannel(ctx, consts.AboutChannelId)
		chAboutChannelDescription <- channel.Description
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		textNewsList, _ := service.Article().MobileHomeTextNewsList(ctx, consts.NewsChannelId)
		chTextNewsList <- textNewsList
		defer close(chTextNewsList)
	}()
	// 最新资讯-图片新闻
	chPicNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		picNewsList, _ := service.Article().MobileHomePicNewsList(ctx, consts.NewsChannelId)
		chPicNewsList <- picNewsList
		defer close(chPicNewsList)
	}()
	// 荣誉资质列表
	chHonorList := make(chan []*model.ImageListItem, 1)
	go func() {
		honerList, _ := service.Image().MobileHomeHonerList(ctx, consts.HonorChannelId)
		chHonorList <- honerList
	}()

	err = service.Response().View(ctx, "/mobile/index/index.html", g.Map{
		"navigation":                    <-chNavigation,                    // 导航
		"aboutChannelUrl":               <-chAboutChannelUrl,               // 关于我们栏目url
		"newsChannelUrl":                <-chNewsChannelUrl,                // 新闻动态栏目url
		"productChannelUrl":             <-chProductChannelUrl,             // 产品展示栏目url
		"contactChannelUrl":             <-chContactChannelUrl,             // 联系我们栏目url
		"guestbookChannelUrl":           <-chGuestbookChannelUrl,           // 在线留言栏目url
		"honorChannelUrl":               <-chHonorChannelUrl,               // 荣誉资质栏目url
		"productFuJianChannelUrl":       <-chProductFuJianChannelUrl,       // 附件栏目url
		"productGangBeiChannelUrl":      <-chProductGangBeiChannelUrl,      // 钢背附件栏目url
		"productJianZhenPianChannelUrl": <-chProductJianZhenPianChannelUrl, // 减震片栏目url
		"adList":                        <-chAdList,                        // banner广告
		"recommendGoodsList":            <-chRecommendGoodsList,            // 产品列表
		"aboutChannelDescription":       <-chAboutChannelDescription,       // 关于我们简介
		"textNewsList":                  <-chTextNewsList,                  // 文字新闻
		"picNewsList":                   <-chPicNewsList,                   // 图片新闻
		"honorList":                     <-chHonorList,                     // 荣誉资质列表
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
