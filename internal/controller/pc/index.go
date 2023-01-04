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
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().Navigation(ctx, 0)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNavigation <- navigation
		defer close(chNavigation)
	}()
	// banner广告
	chAdList := make(chan []*entity.CmsAd, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.PcHomeAdChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页广告耗时"+gconv.String(endTime-startTime)+"毫秒")
		chAdList <- adList
		defer close(chAdList)
	}()
	// 首页50个随机新闻滚动
	chScrollNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		scrollNewsList, _ := service.Article().PcHomeScrollNewsList(ctx, consts.NewsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页新闻滚动耗时"+gconv.String(endTime-startTime)+"毫秒")
		chScrollNewsList <- scrollNewsList
		defer close(chScrollNewsList)
	}()
	// 首页3个随机产品图集
	chRecommendGoodsList := make(chan []*model.ImageListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		recommendGoodsList, _ := service.Image().PcHomeRecommendGoodsList(ctx, consts.GoodsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页随机产品耗时"+gconv.String(endTime-startTime)+"毫秒")
		chRecommendGoodsList <- recommendGoodsList
		defer close(chRecommendGoodsList)
	}()
	// 推荐商品查看更多
	chRecommendGoodsMoreUrl := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		recommendGoodsMoreUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GoodsChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页推荐商品查看更多耗时"+gconv.String(endTime-startTime)+"毫秒")
		chRecommendGoodsMoreUrl <- recommendGoodsMoreUrl
		defer close(chRecommendGoodsMoreUrl)
	}()
	// 关于我们
	chAbout := make(chan *entity.CmsChannel, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		about, _ := service.Channel().HomeAboutChannel(ctx, consts.AboutChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页关于我们耗时"+gconv.String(endTime-startTime)+"毫秒")
		chAbout <- about
		defer close(chAbout)
	}()
	// 关于我们查看更多
	chAboutMoreUrl := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		aboutMoreUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.AboutChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页关于我们查看更多耗时"+gconv.String(endTime-startTime)+"毫秒")
		chAboutMoreUrl <- aboutMoreUrl
		defer close(chAboutMoreUrl)
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		guestbookUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc在线留言栏目url耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGuestbookUrl <- guestbookUrl
		defer close(chGuestbookUrl)
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页产品中心栏目列表耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 产品中心产品分组列表
	chGoodsGroupList := make(chan [][]*model.ImageListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		goodsGroupList, _ := service.Image().PcHomeGoodsGroupList(ctx, consts.GoodsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页产品中心产品分组列表耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGoodsGroupList <- goodsGroupList
		defer close(chGoodsGroupList)
	}()
	// 最新资讯查看更多
	chNewsMoreUrl := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		newsMoreUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.NewsChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页最新资讯查看更多耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNewsMoreUrl <- newsMoreUrl
		defer close(chNewsMoreUrl)
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页最新资讯文字新闻列表"+gconv.String(endTime-startTime)+"毫秒")
		chTextNewsList <- textNewsList
		defer close(chTextNewsList)
	}()
	// 最新资讯-图片新闻
	chPicNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		picNewsList, _ := service.Article().PcHomePicNewsList(ctx, consts.NewsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页最新资讯图片新闻列表"+gconv.String(endTime-startTime)+"毫秒")
		chPicNewsList <- picNewsList
		defer close(chPicNewsList)
	}()
	// 友情链接
	chFriendlyLinkList := make(chan []*entity.CmsFriendlyLink, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		friendlyLinkList, _ := service.FriendlyLink().PcList(ctx)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc首页友情链接耗时"+gconv.String(endTime-startTime)+"毫秒")
		chFriendlyLinkList <- friendlyLinkList
		defer close(chFriendlyLinkList)
	}()

	err = service.Response().View(ctx, "/pc/index/index.html", g.Map{
		"navigation":            <-chNavigation,            // 导航
		"adList":                <-chAdList,                // banner
		"scrollNewsList":        <-chScrollNewsList,        // 新闻滚动
		"recommendGoodsList":    <-chRecommendGoodsList,    // 推荐商品
		"recommendGoodsMoreUrl": <-chRecommendGoodsMoreUrl, // 推荐商品查看更多
		"about":                 <-chAbout,                 // 关于我们
		"aboutMoreUrl":          <-chAboutMoreUrl,          // 关于我们查看更多
		"goodsChannelList":      <-chGoodsChannelList,      // 产品中心栏目列表
		"goodsGroupList":        <-chGoodsGroupList,        // 产品中心产品分组列表
		"newsMoreUrl":           <-chNewsMoreUrl,           // 最新资讯查看更多
		"textNewsList":          <-chTextNewsList,          // 最新资讯-文字新闻
		"picNewsList":           <-chPicNewsList,           // 最新资讯-文字新闻
		"friendlyLinkList":      <-chFriendlyLinkList,      // 友情链接列表
		"guestbookUrl":          <-chGuestbookUrl,          // 在线留言栏目url
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}
