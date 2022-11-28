package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

const (
	// pc首页广告分类id
	pcHomeAdChannelId = 1
)

// Index pc首页
func (c *cIndex) Index(ctx context.Context, req *pc.IndexReq) (res *pc.IndexRes, err error) {
	var chNavigation = make(chan []*model.ChannelPcNavigationListItem, 1)
	var chAdList = make(chan []*entity.CmsAd, 1)
	// 导航栏
	go func() {
		navigation, _ := service.Channel().PcNavigation(ctx)
		chNavigation <- navigation
		close(chNavigation)
		return
	}()
	// banner广告
	go func() {
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, pcHomeAdChannelId)
		chAdList <- adList
		close(chAdList)
		return
	}()
	navigation := <-chNavigation
	adList := <-chAdList
	err = service.Response().View(ctx, "/pc/index/index.html", g.Map{
		"navigation": navigation,
		"adList":     adList,
	})
	if err != nil {
		return nil, err
	}
	return
}
