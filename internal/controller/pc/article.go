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
		startTime := gtime.TimestampMilli()
		articlePageList, err := Article.articlePageList(ctx, req)
		if err != nil {
			return
		}
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc文章列表分页数据耗时"+gconv.String(endTime-startTime)+"毫秒")
		chArticlePageList <- articlePageList
		defer close(chArticlePageList)
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
		goodsChannelList, _ := service.Channel().PcHomeGoodsChannelList(ctx, consts.GoodsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc栏目页产品中心栏目列表耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc栏目页最新资讯文字新闻列表"+gconv.String(endTime-startTime)+"毫秒")
		chTextNewsList <- textNewsList
		defer close(chTextNewsList)
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
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		channelTemplate, _ := service.Channel().PcListTemplate(ctx, channelInfo)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc获取栏目模板"+gconv.String(endTime-startTime)+"毫秒")
		chChannelTemplate <- channelTemplate
	}()
	articlePageList := <-chArticlePageList
	pageInfo := service.PageInfo().PcPageInfo(ctx, articlePageList.Total, articlePageList.Size)
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":      channelInfo,          // 栏目信息
		"navigation":       <-chNavigation,       // 导航
		"tdk":              <-chTDK,              // TDK
		"crumbs":           <-chCrumbs,           // 面包屑导航
		"goodsChannelList": <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":     <-chTextNewsList,     // 最新资讯-文字新闻
		"articlePageList":  articlePageList,      // 文章列表
		"pageInfo":         pageInfo,             // 页码
		"guestbookUrl":     <-chGuestbookUrl,     // 在线留言栏目url
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
	// 更新连击量
	go func() {
		_, err = dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, articleInfo.Id).Increment(dao.CmsArticle.Columns().ClickNum, 1)
	}()
	// 栏目详情
	channelInfo, err := Article.channelInfo(ctx, articleInfo.ChannelId)
	if err != nil {
		return nil, err
	}
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
		pcTDK, _ := service.Channel().PcTDK(ctx, gconv.Uint(articleInfo.ChannelId), gconv.Int64(articleInfo.Id))
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
		goodsChannelList, _ := service.Channel().PcHomeGoodsChannelList(ctx, consts.GoodsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc文章详情页产品中心栏目列表耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGoodsChannelList <- goodsChannelList
		defer close(chGoodsChannelList)
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc文章详情页最新资讯文字新闻列表"+gconv.String(endTime-startTime)+"毫秒")
		chTextNewsList <- textNewsList
		defer close(chTextNewsList)
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
		startTime := gtime.TimestampMilli()
		guestbookUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc在线留言栏目url耗时"+gconv.String(endTime-startTime)+"毫秒")
		chGuestbookUrl <- guestbookUrl
		defer close(chGuestbookUrl)
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		channelTemplate, _ := service.Channel().PcDetailTemplate(ctx, channelInfo)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc获取文章详情模板"+gconv.String(endTime-startTime)+"毫秒")
		chChannelTemplate <- channelTemplate
	}()
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"navigation":       <-chNavigation,       // 导航
		"tdk":              <-chTDK,              // TDK
		"crumbs":           <-chCrumbs,           // 面包屑导航
		"goodsChannelList": <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":     <-chTextNewsList,     // 最新资讯-文字新闻
		"articleInfo":      articleInfo,          // 文章详情
		"prevArticle":      <-chPrevArticle,      // 上一篇
		"nextArticle":      <-chNextArticle,      // 下一篇
		"guestbookUrl":     <-chGuestbookUrl,     // 在线留言栏目url
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
	}
	res = &pc.ArticleListRes{
		List:  articleList,
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
	}
	return
}
