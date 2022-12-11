package article

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcHomeScrollNewsList pc首页滚动新闻
func (s *sArticle) PcHomeScrollNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, channelTid, true)
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

// PcHomeTextNewsList pc首页文字新闻列表
func (s *sArticle) PcHomeTextNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, channelTid, true)
	if err != nil {
		return nil, err
	}
	err = dao.CmsArticle.Ctx(ctx).WhereIn(dao.CmsArticle.Columns().ChannelId, arrAllIds).
		Where(dao.CmsArticle.Columns().Status, 1).
		Where("FIND_IN_SET('p', `flag`) = false").
		OrderAsc(dao.CmsArticle.Columns().Sort).
		OrderDesc(dao.CmsArticle.Columns().Id).
		Limit(5).
		Scan(&out)
	if err != nil {
		return nil, err
	}
	if len(out) > 0 {
		for key, item := range out {
			url, err := service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
			if err != nil {
				return nil, err
			}
			out[key].Router = url
		}
	}
	return
}
func (s *sArticle) PcHomePicNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, channelTid, true)
	if err != nil {
		return nil, err
	}
	err = dao.CmsArticle.Ctx(ctx).WhereIn(dao.CmsArticle.Columns().ChannelId, arrAllIds).
		Where(dao.CmsArticle.Columns().Status, 1).
		Where("FIND_IN_SET('p', `flag`) = true").
		OrderAsc(dao.CmsArticle.Columns().Sort).
		OrderDesc(dao.CmsArticle.Columns().Id).
		Limit(5).
		Scan(&out)
	if err != nil {
		return nil, err
	}
	if len(out) > 0 {
		for key, item := range out {
			url, err := service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
			if err != nil {
				return nil, err
			}
			out[key].Router = url
			out[key].Thumb = service.Util().ImageOrDefaultUrl(item.Thumb)
		}
	}
	return
}
