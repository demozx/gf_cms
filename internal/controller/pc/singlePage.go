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
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, req.Id)
		chNavigation <- navigation
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		defer close(chTDK)
		pcTDK, _ := service.Channel().TDK(ctx, channelInfo.Id, 0)
		chTDK <- pcTDK
	}()
	// 面包屑导航
	chCrumbs := make(chan []*model.ChannelCrumbs, 1)
	go func() {
		defer close(chCrumbs)
		pcCrumbs, _ := service.Channel().Crumbs(ctx, channelInfo.Id)
		chCrumbs <- pcCrumbs
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chGoodsChannelList)
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chTextNewsList)
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		chTextNewsList <- textNewsList
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		defer close(chGuestbookUrl)
		guestbookChannelUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookChannelUrl
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		defer close(chChannelTemplate)
		channelTemplate, _ := service.Channel().PcListTemplate(ctx, channelInfo)
		chChannelTemplate <- channelTemplate
	}()
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":         channelInfo,          // 栏目信息
		"navigation":          <-chNavigation,       // 导航
		"tdk":                 <-chTDK,              // TDK
		"crumbs":              <-chCrumbs,           // 面包屑导航
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":        <-chTextNewsList,     // 最新资讯-文字新闻
		"guestbookChannelUrl": <-chGuestbookUrl,     // 在线留言栏目url
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
