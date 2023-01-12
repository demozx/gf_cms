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
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, 0)
		chNavigation <- navigation
	}()
	// banner广告
	chAdList := make(chan []*entity.CmsAd, 1)
	go func() {
		defer close(chAdList)
		adList, _ := service.AdList().PcHomeListByChannelId(ctx, consts.PcHomeAdChannelId)
		chAdList <- adList
	}()
	// 首页50个随机新闻滚动
	chScrollNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chScrollNewsList)
		scrollNewsList, _ := service.Article().PcHomeScrollNewsList(ctx, consts.NewsChannelId)
		chScrollNewsList <- scrollNewsList
	}()
	// 首页3个随机产品图集
	chRecommendGoodsList := make(chan []*model.ImageListItem, 1)
	go func() {
		defer close(chRecommendGoodsList)
		recommendGoodsList, _ := service.Image().PcHomeRecommendGoodsList(ctx, consts.GoodsChannelId)
		chRecommendGoodsList <- recommendGoodsList
	}()
	// 推荐商品查看更多
	chRecommendGoodsMoreUrl := make(chan string, 1)
	go func() {
		defer close(chRecommendGoodsMoreUrl)
		recommendGoodsMoreUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GoodsChannelId, "")
		chRecommendGoodsMoreUrl <- recommendGoodsMoreUrl
	}()
	// 关于我们
	chAbout := make(chan *entity.CmsChannel, 1)
	go func() {
		defer close(chAbout)
		about, _ := service.Channel().HomeAboutChannel(ctx, consts.AboutChannelId)
		chAbout <- about
	}()
	// 关于我们查看更多
	chAboutMoreUrl := make(chan string, 1)
	go func() {
		defer close(chAboutMoreUrl)
		aboutMoreUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.AboutChannelId, "")
		chAboutMoreUrl <- aboutMoreUrl
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		defer close(chGuestbookUrl)
		guestbookChannelUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookChannelUrl
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chGoodsChannelList)
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
	}()
	// 产品中心产品分组列表
	chGoodsGroupList := make(chan [][]*model.ImageListItem, 1)
	go func() {
		defer close(chGoodsGroupList)
		goodsGroupList, _ := service.Image().PcHomeGoodsGroupList(ctx, consts.GoodsChannelId)
		chGoodsGroupList <- goodsGroupList
	}()
	// 最新资讯查看更多
	chNewsMoreUrl := make(chan string, 1)
	go func() {
		defer close(chNewsMoreUrl)
		newsMoreUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.NewsChannelId, "")
		chNewsMoreUrl <- newsMoreUrl
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chTextNewsList)
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		chTextNewsList <- textNewsList
	}()
	// 最新资讯-图片新闻
	chPicNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chPicNewsList)
		picNewsList, _ := service.Article().PcHomePicNewsList(ctx, consts.NewsChannelId)
		chPicNewsList <- picNewsList
	}()
	// 友情链接
	chFriendlyLinkList := make(chan []*entity.CmsFriendlyLink, 1)
	go func() {
		defer close(chFriendlyLinkList)
		friendlyLinkList, _ := service.FriendlyLink().PcList(ctx)
		chFriendlyLinkList <- friendlyLinkList
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
		"guestbookChannelUrl":   <-chGuestbookUrl,          // 在线留言栏目url
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}
