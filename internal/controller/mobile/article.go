package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

// List pc文章列表
func (c *cArticle) List(ctx context.Context, req *mobile.ArticleListReq) (res *mobile.ArticleListRes, err error) {
	// 栏目详情
	channelInfo, err := Article.channelInfo(ctx, req.ChannelId)
	if err != nil {
		return nil, err
	}
	// 文章列表
	chArticlePageList := make(chan *mobile.ArticleListRes, 1)
	go func() {
		articlePageList, _ := Article.articlePageList(ctx, req)
		chArticlePageList <- articlePageList
		defer close(chArticlePageList)
	}()
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		navigation, _ := service.Channel().Navigation(ctx, gconv.Int(channelInfo.Id))
		chNavigation <- navigation
		defer close(chNavigation)
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		tdk, _ := service.Channel().TDK(ctx, channelInfo.Id, 0)
		chTDK <- tdk
		defer close(chTDK)
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		guestbookUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookUrl
		defer close(chGuestbookUrl)
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		channelTemplate, _ := service.Channel().MobileListTemplate(ctx, channelInfo)
		chChannelTemplate <- channelTemplate
	}()
	articlePageList := <-chArticlePageList
	pageInfo := service.PageInfo().MobilePageInfo(ctx, articlePageList.Total, articlePageList.Size)
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":         channelInfo,          // 栏目信息
		"navigation":          <-chNavigation,       // 导航
		"tdk":                 <-chTDK,              // TDK
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"articlePageList":     articlePageList,      // 文章列表
		"pageInfo":            pageInfo,             // 页码
		"guestbookChannelUrl": <-chGuestbookUrl,     // 在线留言栏目url
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// Detail mobile文章详情页面
func (c *cArticle) Detail(ctx context.Context, req *mobile.ArticleDetailReq) (res *mobile.ArticleDetailRes, err error) {
	// 文章详情
	var articleInfo *model.ArticleWithBody
	err = dao.CmsArticle.Ctx(ctx).
		Where(dao.CmsArticle.Columns().Id, req.Id).
		Where(dao.CmsArticle.Columns().Status, 1).
		With(model.ArticleBodyItem{}).
		Scan(&articleInfo)
	if err != nil {
		return nil, err
	}
	if articleInfo == nil {
		service.Response().Status404(ctx)
	}
	articleInfo.ClickNum++
	// 更新点击量
	go func() {
		_, err = dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, articleInfo.Id).Increment(dao.CmsArticle.Columns().ClickNum, 1)
	}()
	// 栏目详情
	channelInfo, err := Article.channelInfo(ctx, articleInfo.ChannelId)
	if err != nil {
		return nil, err
	}
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		navigation, _ := service.Channel().Navigation(ctx, gconv.Int(channelInfo.Id))
		chNavigation <- navigation
		defer close(chNavigation)
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		tdk, _ := service.Channel().TDK(ctx, gconv.Uint(articleInfo.ChannelId), gconv.Int64(articleInfo.Id))
		chTDK <- tdk
		defer close(chTDK)
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 上一篇
	chPrevArticle := make(chan *model.ArticleLink, 1)
	go func() {
		prevArticle, _ := service.Article().PcPrevArticle(ctx, articleInfo.ChannelId, articleInfo.Id)
		chPrevArticle <- prevArticle
	}()
	// 下一篇
	chNextArticle := make(chan *model.ArticleLink, 1)
	go func() {
		nextArticle, _ := service.Article().PcNextArticle(ctx, articleInfo.ChannelId, articleInfo.Id)
		chNextArticle <- nextArticle
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		guestbookUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookUrl
		defer close(chGuestbookUrl)
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		channelTemplate, _ := service.Channel().MobileDetailTemplate(ctx, channelInfo)
		chChannelTemplate <- channelTemplate
	}()
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"navigation":          <-chNavigation,       // 导航
		"tdk":                 <-chTDK,              // TDK
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"articleInfo":         articleInfo,          // 文章详情
		"prevArticle":         <-chPrevArticle,      // 上一篇
		"nextArticle":         <-chNextArticle,      // 下一篇
		"guestbookChannelUrl": <-chGuestbookUrl,     // 在线留言栏目url
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// 栏目详情
func (c *cArticle) channelInfo(ctx context.Context, channelId int) (out *entity.CmsChannel, err error) {
	err = dao.CmsChannel.Ctx(ctx).
		Where(dao.CmsChannel.Columns().Id, channelId).
		Where(dao.CmsChannel.Columns().Status, 1).
		Where(dao.CmsChannel.Columns().Type, 1).
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

// 获取文章列表分页数据
func (c *cArticle) articlePageList(ctx context.Context, in *mobile.ArticleListReq) (res *mobile.ArticleListRes, err error) {
	// 当前栏目的所有级别的子栏目id们加自己
	childChannelIds, err := service.Channel().GetChildIds(ctx, in.ChannelId, true)
	if err != nil {
		return nil, err
	}
	m := dao.CmsArticle.Ctx(ctx).
		WhereIn(dao.CmsArticle.Columns().ChannelId, childChannelIds).
		Where(dao.CmsArticle.Columns().Status, 1)
	count, err := m.Count()
	if err != nil {
		return
	}
	var articleList []*model.ArticleListItem
	err = m.OrderAsc(dao.CmsArticle.Columns().Sort).
		OrderDesc(dao.CmsArticle.Columns().Id).
		Page(in.Page, in.Size).
		Scan(&articleList)
	if err != nil {
		return nil, err
	}
	for key, item := range articleList {
		url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
		if err != nil {
			return nil, err
		}
		articleList[key].Router = url
	}
	res = &mobile.ArticleListRes{
		List:  articleList,
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
	}
	return
}
