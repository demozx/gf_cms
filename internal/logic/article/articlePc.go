package article

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcHomeScrollNewsBelongChannelId pc首页滚动新闻
func (s *sArticle) PcHomeScrollNewsBelongChannelId(ctx context.Context, belongChannelId int) (out []*model.ArticleListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, belongChannelId, true)
	if err != nil {
		return nil, err
	}
	var scrollNewsList []*model.ArticleListItem
	err = dao.CmsArticle.Ctx(ctx).
		WhereIn(dao.CmsArticle.Columns().ChannelId, arrAllIds).
		Where(dao.CmsArticle.Columns().Status, 1).
		OrderRandom().
		Limit(50).
		Fields(dao.CmsArticle.Columns().Id, dao.CmsArticle.Columns().Title).
		Scan(&scrollNewsList)
	if err != nil {
		return nil, err
	}
	for key, item := range scrollNewsList {
		detailUrl, _ := service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
		scrollNewsList[key].Router = detailUrl
	}
	out = scrollNewsList
	return
}
