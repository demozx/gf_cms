package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Image = cImage{}
)

type cImage struct{}

// List pc图集列表
func (c *cImage) List(ctx context.Context, req *pc.ImageListReq) (res *pc.ImageListRes, err error) {
	// 栏目详情
	channelInfo, err := Article.channelInfo(ctx, req.ChannelId)
	if err != nil {
		return nil, err
	}
	// 图集列表
	chImagePageList := make(chan *pc.ImageListRes, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		imagePageList, err := Image.imagePageList(ctx, req)
		if err != nil {
			return
		}
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc图集列表分页数据耗时"+gconv.String(endTime-startTime)+"毫秒")
		chImagePageList <- imagePageList
		defer close(chImagePageList)
	}()
	// 导航栏
	chNavigation := make(chan []*model.ChannelPcNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().PcNavigation(ctx, gconv.Int(channelInfo.Id))
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
	imagePageList := <-chImagePageList
	pageInfo := service.PageInfo().PcPageInfo(ctx, imagePageList.Total, imagePageList.Size)
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":      channelInfo,          // 栏目信息
		"navigation":       <-chNavigation,       // 导航
		"tdk":              <-chTDK,              // TDK
		"crumbs":           <-chCrumbs,           // 面包屑导航
		"goodsChannelList": <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":     <-chTextNewsList,     // 最新资讯-文字新闻
		"imagePageList":    imagePageList,        // 图集列表
		"pageInfo":         pageInfo,             // 页码
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// 获取文章列表分页数据
func (c *cImage) imagePageList(ctx context.Context, in *pc.ImageListReq) (res *pc.ImageListRes, err error) {
	// 当前栏目的所有级别的子栏目id们加自己
	childChannelIds, err := service.Channel().GetChildIds(ctx, in.ChannelId, true)
	if err != nil {
		return nil, err
	}
	m := dao.CmsImage.Ctx(ctx).
		WhereIn(dao.CmsImage.Columns().ChannelId, childChannelIds).
		Where(dao.CmsImage.Columns().Status, 1)
	count, err := m.Count()
	if err != nil {
		return
	}
	var imageList []*model.ImageListItem
	err = m.OrderAsc(dao.CmsImage.Columns().Sort).
		OrderDesc(dao.CmsImage.Columns().Id).
		Page(in.Page, in.Size).
		Scan(&imageList)
	if err != nil {
		return nil, err
	}
	for key, item := range imageList {
		url, err := service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelImage, gconv.Int(item.Id))
		if err != nil {
			return nil, err
		}
		imageList[key].Router = url
		imageList[key], err = service.Image().BuildThumb(ctx, item)
		if err != nil {
			return nil, err
		}
	}
	res = &pc.ImageListRes{
		List:  imageList,
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
	}
	return
}
