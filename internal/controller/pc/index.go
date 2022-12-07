package pc

import (
	"context"
	"gf_cms/api/pc"
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
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().PcNavigation(ctx)
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNavigation <- navigation
		close(chNavigation)
	}()
	// banner广告
	go func() {
		startTime := gtime.TimestampMilli()
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.PcHomeAdChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc首页广告耗时"+gconv.String(endTime-startTime)+"毫秒")
		chAdList <- adList
		close(chAdList)
	}()
	// 首页50个随机新闻滚动
	go func() {
		startTime := gtime.TimestampMilli()
		scrollNewsList, _ := service.Article().PcHomeScrollNewsBelongChannelId(ctx, consts.PcHomeScrollNewsBelongChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc首页新闻滚动耗时"+gconv.String(endTime-startTime)+"毫秒")
		chScrollNewsList <- scrollNewsList
		close(chScrollNewsList)
	}()
	// 首页3个随机产品图集
	go func() {
		startTime := gtime.TimestampMilli()
		recommendGoodsList, _ := service.Image().PcHomeRecommendGoodsList(ctx, consts.PcHomeRecommendGoodsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc首页随机产品耗时"+gconv.String(endTime-startTime)+"毫秒")
		chRecommendGoodsList <- recommendGoodsList
		close(chRecommendGoodsList)
	}()
	// 推荐商品查看更多
	go func() {
		startTime := gtime.TimestampMilli()
		recommendGoodsMoreUrl, _ := service.GenUrl().PcChannelUrl(ctx, consts.PcHomeRecommendGoodsChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc首页推荐商品查看更多耗时"+gconv.String(endTime-startTime)+"毫秒")
		chRecommendGoodsMoreUrl <- recommendGoodsMoreUrl
		close(chRecommendGoodsMoreUrl)
	}()
	// 关于我们
	go func() {
		startTime := gtime.TimestampMilli()
		about, _ := service.Channel().PcHomeAboutChannel(ctx, consts.AbortChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc首页关于我们耗时"+gconv.String(endTime-startTime)+"毫秒")
		chAbout <- about
		close(chAbout)
	}()
	// 关于我们查看更多
	go func() {
		startTime := gtime.TimestampMilli()
		aboutMoreUrl, _ := service.GenUrl().PcChannelUrl(ctx, consts.AbortChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Info(ctx, "pc首页关于我们查看更多耗时"+gconv.String(endTime-startTime)+"毫秒")
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
