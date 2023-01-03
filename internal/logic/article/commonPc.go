package article

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcPrevArticle 上一篇文章
func (s *sArticle) PcPrevArticle(ctx context.Context, channelId int, articleId uint64) (out *model.ArticleLink, err error) {
	var prevArticle *entity.CmsArticle
	err = dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().ChannelId, channelId).Where(dao.CmsArticle.Columns().Status, 1).WhereLT(dao.CmsArticle.Columns().Id, articleId).OrderAsc(dao.CmsArticle.Columns().Sort).OrderDesc(dao.CmsArticle.Columns().Id).Scan(&prevArticle)
	if err != nil {
		return
	}
	if prevArticle == nil {
		out = &model.ArticleLink{Title: "无", Router: "javascript:;"}
	} else {
		url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(prevArticle.Id))
		if err != nil {
			return nil, err
		}
		out = &model.ArticleLink{Title: prevArticle.Title, Router: url}
	}
	return
}

// PcNextArticle 下一篇文章
func (s *sArticle) PcNextArticle(ctx context.Context, channelId int, articleId uint64) (out *model.ArticleLink, err error) {
	var nextArticle *entity.CmsArticle
	err = dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().ChannelId, channelId).Where(dao.CmsArticle.Columns().Status, 1).WhereGT(dao.CmsArticle.Columns().Id, articleId).OrderAsc(dao.CmsArticle.Columns().Sort).OrderAsc(dao.CmsArticle.Columns().Id).Scan(&nextArticle)
	if err != nil {
		return
	}
	if nextArticle == nil {
		out = &model.ArticleLink{Title: "无", Router: "javascript:;"}
	} else {
		url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelArticle, gconv.Int(nextArticle.Id))
		if err != nil {
			return nil, err
		}
		out = &model.ArticleLink{Title: nextArticle.Title, Router: url}
	}
	return
}
