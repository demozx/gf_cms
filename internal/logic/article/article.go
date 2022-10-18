package article

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
)

type (
	sArticle struct{}
)

var (
	insArticle = sArticle{}
)

func init() {
	service.RegisterArticle(New())
}

func New() *sArticle {
	return &sArticle{}
}

func Article() *sArticle {
	return &insArticle
}

func (s *sArticle) BackendArticleGetList(ctx context.Context, in *model.ArticleGetListInPut) (out *model.ArticleGetListOutPut, err error) {
	m := dao.CmsArticle.Ctx(ctx).As("article").OrderAsc("article.sort").OrderDesc("article.id")
	out = &model.ArticleGetListOutPut{
		Page: in.Page,
		Size: in.Size,
	}
	if in.ChannelId > 0 {
		m = m.Where("article.channel_id", in.ChannelId)
	}
	if in.StartAt != "" && in.EndAt != "" {
		m = m.WhereGTE("article.created_at", in.StartAt).WhereLT("article.created_at", in.EndAt)
	}
	if in.Keyword != "" {
		m = m.WhereLike("article.title", "%"+in.Keyword+"%")
	}
	listModel := m.LeftJoin(dao.CmsChannel.Table(), "channel", "channel.id=article.channel_id").
		Fields("article.*, channel.name channel_name").
		Page(in.Page, in.Size)

	var list []*model.ArticleListItem
	err = listModel.Scan(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sArticle) Sort(ctx context.Context, in []*model.ArticleSortMap) (out interface{}, err error) {
	update, err := dao.CmsArticle.Ctx(ctx).Save(in)
	if err != nil {
		return false, err
	}
	return update, nil
}
