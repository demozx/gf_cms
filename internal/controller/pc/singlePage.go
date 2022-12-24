package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	SinglePage = cSinglePage{}
)

type cSinglePage struct{}

// Detail pc单页
func (c *cSinglePage) Detail(ctx context.Context, req *pc.SinglePageReq) (res *pc.SinglePageRes, err error) {
	// 内容详情
	channelInfo, err := SinglePage.channelInfo(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 导航栏
	chNavigation := make(chan []*model.ChannelPcNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().PcNavigation(ctx, req.Id)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNavigation <- navigation
		defer close(chNavigation)
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		pcTDK, _ := service.Channel().PcTDK(ctx, channelInfo.Id, 0)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pcTDK耗时"+gconv.String(endTime-startTime)+"毫秒")
		chTDK <- pcTDK
		defer close(chTDK)
	}()
	// 面包屑导航
	chCrumbs := make(chan []*model.ChannelCrumbs, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		pcCrumbs, _ := service.Channel().PcCrumbs(ctx, channelInfo.Id)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc面包屑导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chCrumbs <- pcCrumbs
		defer close(chCrumbs)
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelPcNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		goodsChannelList, _ := service.Channel().PcHomeGoodsChannelList(ctx, consts.GoodsChannelTid)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc栏目页产品中心栏目列表耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelTid)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc栏目页最新资讯文字新闻列表"+gconv.String(endTime-startTime)+"毫秒")
		chTextNewsList <- textNewsList
		defer close(chTextNewsList)
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		channelTemplate, _ := service.Channel().PcListTemplate(ctx, channelInfo)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc获取栏目模板"+gconv.String(endTime-startTime)+"毫秒")
		chChannelTemplate <- channelTemplate
	}()
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":      channelInfo,          // 栏目信息
		"navigation":       <-chNavigation,       // 导航
		"tdk":              <-chTDK,              // TDK
		"crumbs":           <-chCrumbs,           // 面包屑导航
		"goodsChannelList": <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":     <-chTextNewsList,     // 最新资讯-文字新闻
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// 栏目详情
func (c *cSinglePage) channelInfo(ctx context.Context, channelId int) (out *entity.CmsChannel, err error) {
	err = dao.CmsChannel.Ctx(ctx).
		Where(dao.CmsChannel.Columns().Id, channelId).
		Where(dao.CmsChannel.Columns().Status, 1).
		Where(dao.CmsChannel.Columns().Type, 2).
		Scan(&out)
	if err != nil {
		return
	}
	// 栏目不存在，展示404
	if out == nil {
		service.Response().Status404(ctx)
	}
	return
}
