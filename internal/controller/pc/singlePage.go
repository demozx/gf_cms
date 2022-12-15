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
	var channelInfo *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, req.Id).Where(dao.CmsChannel.Columns().Status, 1).Scan(&channelInfo)
	if err != nil {
		return
	}
	// 栏目不存在，展示404
	if channelInfo == nil {
		service.Response().Pc404(ctx)
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
	// 面包屑导航
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelPcNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		goodsChannelList, _ := service.Channel().PcHomeGoodsChannelList(ctx, consts.GoodsChannelTid)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页产品中心栏目列表耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelTid)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页最新资讯文字新闻列表"+gconv.String(endTime-startTime)+"毫秒")
		chTextNewsList <- textNewsList
		defer close(chTextNewsList)
	}()
	err = service.Response().View(ctx, "/pc/single_page/detail.html", g.Map{
		"navigation":       <-chNavigation,       // 导航
		"goodsChannelList": <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":     <-chTextNewsList,     // 最新资讯-文字新闻
	})
	if err != nil {
		return nil, err
	}
	return
}
