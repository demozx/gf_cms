package article

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"sync"
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

func (s *sArticle) Flag(ctx context.Context, ids []int, flagType string) (out interface{}, err error) {
	if len(ids) == 1 {
		_, err = Article().singleFlag(ctx, ids[0], flagType, "auto")
		if err != nil {
			return nil, err
		}
	} else {
		var wg sync.WaitGroup
		for _, id := range ids {
			wg.Add(1)
			go func(id int) {
				_, err = Article().singleFlag(ctx, id, flagType, "open")
				defer wg.Done()
				if err != nil {
					return
				}
			}(id)
		}
		wg.Wait()
	}
	return
}

func (s *sArticle) Status(ctx context.Context, ids []int) (out interface{}, err error) {
	if len(ids) == 1 {
		_, err = Article().singleStatus(ctx, ids[0], "auto")
		if err != nil {
			return nil, err
		}
	} else {
		var wg sync.WaitGroup
		for _, id := range ids {
			wg.Add(1)
			go func(id int) {
				_, err = Article().singleStatus(ctx, id, "open")
				defer wg.Done()
				if err != nil {
					return
				}
			}(id)
		}
		wg.Wait()
	}
	return
}

func (s *sArticle) Delete(ctx context.Context, ids []int) (out interface{}, err error) {
	mArticle := dao.CmsArticle.Ctx(ctx).WhereIn(dao.CmsArticle.Columns().Id, ids)
	mBody := dao.CmsArticleBody.Ctx(ctx).WhereIn(dao.CmsArticleBody.Columns().ArticleId, ids)
	if service.Util().GetSetting("recycle_bin") == "1" {
		_, err = mArticle.Delete()
		_, err = mBody.Delete()
	} else {
		_, err = mArticle.Unscoped().Delete()
		_, err = mBody.Unscoped().Delete()
	}

	if err != nil {
		return nil, err
	}
	return
}

func (s *sArticle) singleFlag(ctx context.Context, id int, flagType string, targetType string) (out interface{}, err error) {
	m := dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, id)
	var article *entity.CmsArticle
	err = m.Scan(&article)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, gerror.New("数据不存在")
	}
	split := gstr.SplitAndTrim(article.Flag, ",")
	if targetType == "auto" {
		if gstr.InArray(split, flagType) {
			for index, value := range split {
				if value == flagType {
					split = append(split[:index], split[index+1:]...)
					break
				}
			}
		} else {
			split = append(split, flagType)
		}
	} else if targetType == "open" {
		if !gstr.InArray(split, flagType) {
			split = append(split, flagType)
		}
	} else if targetType == "close" {
		if gstr.InArray(split, flagType) {
			for index, value := range split {
				if value == flagType {
					split = append(split[:index], split[index+1:]...)
					break
				}
			}
		}
	} else {
		return nil, gerror.New("操作目的错误")
	}
	flag := gstr.Implode(",", split)
	_, err = m.Data(g.Map{
		dao.CmsArticle.Columns().Flag: flag,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sArticle) singleStatus(ctx context.Context, id int, targetType string) (out interface{}, err error) {
	m := dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, id)
	var article *entity.CmsArticle
	err = m.Scan(&article)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, gerror.New("数据不存在")
	}
	status := 0
	if targetType == "auto" {
		if article.Status == 0 {
			status = 1
		}
	} else if targetType == "open" {
		status = 1
	} else if targetType == "close" {
		status = 0
	} else {
		return nil, gerror.New("操作目的错误")
	}
	_, err = m.Data(g.Map{
		dao.CmsArticle.Columns().Status: status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}
