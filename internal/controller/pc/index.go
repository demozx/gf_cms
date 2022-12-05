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
	// 导航栏
	go func() {
		navigation, _ := service.Channel().PcNavigation(ctx)
		chNavigation <- navigation
		close(chNavigation)
		return
	}()
	// banner广告
	go func() {
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.PcHomeAdChannelId)
		chAdList <- adList
		close(chAdList)
		return
	}()
	// 首页新闻滚动
	go func() {
		scrollNewsList, _ := service.Article().PcHomeScrollNewsBelongChannelId(ctx, consts.PcHomeScrollNewsBelongChannelId)
		chScrollNewsList <- scrollNewsList
		close(chScrollNewsList)
		return
	}()
	// 首页3个随机图集

	err = service.Response().View(ctx, "/pc/index/index.html", g.Map{
		"navigation":     <-chNavigation,
		"adList":         <-chAdList,
		"scrollNewsList": <-chScrollNewsList,
	})
	if err != nil {
		return nil, err
	}
	return
}
