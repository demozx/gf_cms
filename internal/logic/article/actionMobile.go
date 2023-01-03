package article

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// MobileHomeTextNewsList 移动首页文字新闻列表
func (s *sArticle) MobileHomeTextNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, channelTid, true)
	if err != nil {
		return nil, err
	}
	err = dao.CmsArticle.Ctx(ctx).WhereIn(dao.CmsArticle.Columns().ChannelId, arrAllIds).
		Where(dao.CmsArticle.Columns().Status, 1).
		Where("FIND_IN_SET('p', `flag`) = false").
		OrderAsc(dao.CmsArticle.Columns().Sort).
		OrderDesc(dao.CmsArticle.Columns().Id).
		Limit(3).
		Scan(&out)
	if err != nil {
		return nil, err
	}
	if len(out) > 0 {
		for key, item := range out {
			url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
			if err != nil {
				return nil, err
			}
			out[key].Router = url
		}
	}
	return
}

func (s *sArticle) MobileHomePicNewsList(ctx context.Context, channelTid int) (out []*model.ArticleListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, channelTid, true)
	if err != nil {
		return nil, err
	}
	err = dao.CmsArticle.Ctx(ctx).WhereIn(dao.CmsArticle.Columns().ChannelId, arrAllIds).
		Where(dao.CmsArticle.Columns().Status, 1).
		Where("FIND_IN_SET('p', `flag`) = true").
		OrderAsc(dao.CmsArticle.Columns().Sort).
		OrderDesc(dao.CmsArticle.Columns().Id).
		Limit(1).
		Scan(&out)
	if err != nil {
		return nil, err
	}
	if len(out) > 0 {
		for key, item := range out {
			url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(item.Id))
			if err != nil {
				return nil, err
			}
			out[key].Router = url
			out[key].Thumb = service.Util().ImageOrDefaultUrl(item.Thumb)
		}
	}
	return
}
