package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Search = cSearch{}
)

type cSearch struct{}

func (c *cSearch) Index(ctx context.Context, req *mobile.SearchReq) (res *mobile.SearchRes, err error) {
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		navigation, _ := service.Channel().Navigation(ctx, 0)
		chNavigation <- navigation
		defer close(chNavigation)
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
		guestbookChannelUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookChannelUrl
		defer close(chGuestbookUrl)
	}()
	// 搜索结果
	list, err := Search.searchArticleList(ctx, req)
	if err != nil {
		return nil, err
	}
	total := 0
	size := 0
	if list != nil {
		total = list.Total
		size = list.Size
	}
	pageInfo := service.PageInfo().MobilePageInfo(ctx, total, size)
	err = service.Response().View(ctx, "/mobile/search/index.html", g.Map{
		"navigation":          <-chNavigation,       // 导航
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"guestbookChannelUrl": <-chGuestbookUrl,     // 在线留言栏目url
		"keyword":             req.Keyword,          // 搜索关键词
		"list":                list,                 // 搜索结果列表
		"pageInfo":            pageInfo,             // 分页
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// 搜索结果列表
func (c *cSearch) searchArticleList(ctx context.Context, in *mobile.SearchReq) (res *mobile.SearchRes, err error) {
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
	res = &mobile.SearchRes{
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
		List:  list,
	}
	return
}
