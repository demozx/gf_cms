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
	chNavigation := make(chan []*model.ChannelPcNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().PcNavigation(ctx, 0)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "mobile导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNavigation <- navigation
		defer close(chNavigation)
	}()
	// 公司简介栏目url
	chAboutChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.AboutChannelTid, "")
		chAboutChannelUrl <- url
		defer close(chAboutChannelUrl)
	}()
	// 新闻动态栏目url
	chNewsChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.NewsChannelTid, "")
		chNewsChannelUrl <- url
		defer close(chNewsChannelUrl)
	}()
	// 产品展示栏目url
	chProductChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.ProductChannelTid, "")
		chProductChannelUrl <- url
		defer close(chProductChannelUrl)
	}()
	// 联系我们栏目url
	chContactChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.ContactChannelTid, "")
		chContactChannelUrl <- url
		defer close(chContactChannelUrl)
	}()
	// 附件栏目url
	chProductFuJianChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.ProductFujianChannelId, "")
		chProductFuJianChannelUrl <- url
		defer close(chProductFuJianChannelUrl)
	}()
	// 钢背附件栏目url
	chProductGangBeiChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.ProductGangBeiChannelId, "")
		chProductGangBeiChannelUrl <- url
		defer close(chProductGangBeiChannelUrl)
	}()
	// 减震片栏目url
	chProductJianZhenPianChannelUrl := make(chan string, 1)
	go func() {
		url, _ := service.GenUrl().PcChannelUrl(ctx, consts.ProductJianZhenPianChannelId, "")
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

	err = service.Response().View(ctx, "/mobile/index/index.html", g.Map{
		"navigation":                    <-chNavigation,                    // 导航
		"aboutChannelUrl":               <-chAboutChannelUrl,               // 公司简介栏目url
		"newsChannelUrl":                <-chNewsChannelUrl,                // 新闻动态栏目url
		"productChannelUrl":             <-chProductChannelUrl,             // 产品展示栏目url
		"contactChannelUrl":             <-chContactChannelUrl,             // 联系我们栏目url
		"productFuJianChannelUrl":       <-chProductFuJianChannelUrl,       // 附件栏目url
		"productGangBeiChannelUrl":      <-chProductGangBeiChannelUrl,      // 钢背附件栏目url
		"productJianZhenPianChannelUrl": <-chProductJianZhenPianChannelUrl, // 减震片栏目url
		"adList":                        <-chAdList,                        // banner广告
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
