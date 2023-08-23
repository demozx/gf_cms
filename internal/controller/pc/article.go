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
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

// List pc文章列表
func (c *cArticle) List(ctx context.Context, req *pc.ArticleListReq) (res *pc.ArticleListRes, err error) {
	// 栏目详情
	channelInfo, err := Article.channelInfo(ctx, req.ChannelId)
	if err != nil {
		return nil, err
	}
	// 文章列表
	chArticlePageList := make(chan *pc.ArticleListRes, 1)
	go func() {
		defer close(chArticlePageList)
		articlePageList, _ := Article.articlePageList(ctx, req)
		chArticlePageList <- articlePageList
	}()
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, gconv.Int(channelInfo.Id))
		chNavigation <- navigation
	}()
	navigation := <-chNavigation
	// 子栏目
	chChildrenNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chChildrenNavigation)
		childrenNavigation, _ := service.Channel().ChildrenNavigation(ctx, navigation, gconv.Int(channelInfo.Id))
		chChildrenNavigation <- childrenNavigation
	}()
	childrenNavigation := <-chChildrenNavigation
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		defer close(chTDK)
		tdk, _ := service.Channel().TDK(ctx, channelInfo.Id, 0)
		chTDK <- tdk
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
	articlePageList := <-chArticlePageList
	pageInfo := service.PageInfo().PcPageInfo(ctx, articlePageList.Total, articlePageList.Size)
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":         channelInfo,          // 栏目信息
		"navigation":          navigation,           // 导航
		"childrenNavigation":  childrenNavigation,   // 子栏目
		"tdk":                 <-chTDK,              // TDK
		"crumbs":              <-chCrumbs,           // 面包屑导航
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":        <-chTextNewsList,     // 最新资讯-文字新闻
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

// Detail pc文章详情页面
func (c *cArticle) Detail(ctx context.Context, req *pc.ArticleDetailReq) (res *pc.ArticleDetailRes, err error) {
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
		ctx := context.Background()
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
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, gconv.Int(channelInfo.Id))
		chNavigation <- navigation
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		defer close(chTDK)
		tdk, _ := service.Channel().TDK(ctx, gconv.Uint(articleInfo.ChannelId), gconv.Int64(articleInfo.Id))
		chTDK <- tdk
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
	// 上一篇
	chPrevArticle := make(chan *model.ArticleLink, 1)
	go func() {
		defer close(chPrevArticle)
		prevArticle, _ := service.Article().PcPrevArticle(ctx, articleInfo.ChannelId, articleInfo.Id)
		chPrevArticle <- prevArticle
	}()
	// 下一篇
	chNextArticle := make(chan *model.ArticleLink, 1)
	go func() {
		defer close(chNextArticle)
		nextArticle, _ := service.Article().PcNextArticle(ctx, articleInfo.ChannelId, articleInfo.Id)
		chNextArticle <- nextArticle
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
		channelTemplate, _ := service.Channel().PcDetailTemplate(ctx, channelInfo)
		chChannelTemplate <- channelTemplate
	}()
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"navigation":          <-chNavigation,       // 导航
		"tdk":                 <-chTDK,              // TDK
		"crumbs":              <-chCrumbs,           // 面包屑导航
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":        <-chTextNewsList,     // 最新资讯-文字新闻
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
func (c *cArticle) articlePageList(ctx context.Context, in *pc.ArticleListReq) (res *pc.ArticleListRes, err error) {
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
		articleList[key].Thumb = service.Util().ImageOrDefaultUrl(item.Thumb)
	}
	res = &pc.ArticleListRes{
		List:  articleList,
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
	}
	return
}
