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
	Search = cSearch{}
)

type cSearch struct{}

func (c *cSearch) Index(ctx context.Context, req *pc.SearchReq) (res *pc.SearchRes, err error) {
	// 导航栏
	chNavigation := make(chan []*model.ChannelPcNavigationListItem, 1)
	go func() {
		startTime := gtime.TimestampMilli()
		navigation, _ := service.Channel().PcNavigation(ctx, 0)
		endTime := gtime.TimestampMilli()
		g.Log().Async().Info(ctx, "pc导航耗时"+gconv.String(endTime-startTime)+"毫秒")
		chNavigation <- navigation
		defer close(chNavigation)
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
	// 搜索结果
	list, err := Search.searchArticleList(ctx, req)
	if err != nil {
		return nil, err
	}
	pageInfo := service.PageInfo().PcPageInfo(ctx, list.Total, list.Size)
	err = service.Response().View(ctx, "/pc/search/index.html", g.Map{
		"navigation":       <-chNavigation,       // 导航
		"goodsChannelList": <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":     <-chTextNewsList,     // 最新资讯-文字新闻
		"guestbookUrl":     <-chGuestbookUrl,     // 在线留言栏目url
		"keyword":          req.Keyword,          // 搜索关键词
		"list":             list,                 // 搜索结果列表
		"pageInfo":         pageInfo,             // 分页
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// 搜索结果列表
func (c *cSearch) searchArticleList(ctx context.Context, in *pc.SearchReq) (res *pc.SearchRes, err error) {
	var list []*model.ArticleListItem
	m := dao.CmsArticle.Ctx(ctx).
		Where(dao.CmsArticle.Columns().Status, 1).
		WhereLike(dao.CmsArticle.Columns().Title, "%"+in.Keyword+"%").
		WhereOrLike(dao.CmsArticle.Columns().Description, "%"+in.Keyword+"%")
	count, err := m.Count()
	if err != nil {
		return
	}
	if count == 0 {
		return
	}
	err = m.Page(in.Page, in.Size).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	for key, item := range list {
		url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
		if err != nil {
			return nil, err
		}
		list[key].Router = url
	}
	res = &pc.SearchRes{
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
		List:  list,
	}
	return
}
