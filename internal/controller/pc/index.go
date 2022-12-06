package pc

import (
	"context"
	"gf_cms/api/pc"
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

// Index pc首页
func (c *cIndex) Index(ctx context.Context, req *pc.IndexReq) (res *pc.IndexRes, err error) {
	var chNavigation = make(chan []*model.ChannelPcNavigationListItem, 1)
	var chAdList = make(chan []*entity.CmsAd, 1)
	var chScrollNewsList = make(chan []*model.ArticleListItem, 1)
	var chRecommendGoodsList = make(chan []*model.ImageListItem, 1)
	var chRecommendGoodsMoreUrl = make(chan string, 1)
	var chAbout = make(chan *entity.CmsChannel, 1)
	var chAboutMoreUrl = make(chan string, 1)
	// 导航栏
	go func() {
		navigation, _ := service.Channel().PcNavigation(ctx)
		chNavigation <- navigation
		close(chNavigation)
	}()
	// banner广告
	go func() {
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.PcHomeAdChannelId)
		chAdList <- adList
		close(chAdList)
	}()
	// 首页50个随机新闻滚动
	go func() {
		scrollNewsList, _ := service.Article().PcHomeScrollNewsBelongChannelId(ctx, consts.PcHomeScrollNewsBelongChannelId)
		chScrollNewsList <- scrollNewsList
		close(chScrollNewsList)
	}()
	// 首页3个随机产品图集
	go func() {
		recommendGoodsList, _ := service.Image().PcHomeRecommendGoodsList(ctx, consts.PcHomeRecommendGoodsChannelId)
		chRecommendGoodsList <- recommendGoodsList
		close(chRecommendGoodsList)
	}()
	// 推荐商品查看更多
	go func() {
		recommendGoodsMoreUrl, _ := service.GenUrl().PcChannelUrl(ctx, consts.PcHomeRecommendGoodsChannelId, "")
		chRecommendGoodsMoreUrl <- recommendGoodsMoreUrl
		close(chRecommendGoodsMoreUrl)
	}()
	// 关于我们
	go func() {
		about, _ := service.Channel().PcHomeAboutChannel(ctx, consts.AbortChannelId)
		chAbout <- about
		close(chAbout)
	}()
	// 关于我们查看更多
	go func() {
		aboutMoreUrl, _ := service.GenUrl().PcChannelUrl(ctx, consts.AbortChannelId, "")
		chAboutMoreUrl <- aboutMoreUrl
		close(chAboutMoreUrl)
	}()
	err = service.Response().View(ctx, "/pc/index/index.html", g.Map{
		"navigation":            <-chNavigation,            // 导航
		"adList":                <-chAdList,                // banner
		"scrollNewsList":        <-chScrollNewsList,        // 新闻滚动
		"recommendGoodsList":    <-chRecommendGoodsList,    // 推荐商品
		"recommendGoodsMoreUrl": <-chRecommendGoodsMoreUrl, // 推荐商品查看更多
		"about":                 <-chAbout,                 // 关于我们
		"aboutMoreUrl":          <-chAboutMoreUrl,          // 关于我们查看更多
	})
	if err != nil {
		return nil, err
	}
	return
}
