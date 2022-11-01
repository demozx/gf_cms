package article

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"regexp"
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
	for i, item := range out.List {
		out.List[i].Thumb = service.Util().ImageOrDefaultUrl(item.Thumb)
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

func (s *sArticle) Move(ctx context.Context, channelId int, ids []string) (out interface{}, err error) {
	if channelId <= 0 {
		return nil, gerror.New("频道ID错误")
	}
	if len(ids) == 0 {
		return nil, gerror.New("要移动的文章不能为空")
	}
	_, err = dao.CmsArticle.Ctx(ctx).WhereIn(dao.CmsArticle.Columns().Id, ids).Data(g.Map{
		dao.CmsArticle.Columns().ChannelId: channelId,
	}).Update()
	if err != nil {
		return nil, err
	}
	_, err = dao.CmsArticleBody.Ctx(ctx).WhereIn(dao.CmsArticleBody.Columns().ArticleId, ids).Data(g.Map{
		dao.CmsArticle.Columns().ChannelId: channelId,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sArticle) Add(ctx context.Context, in *backendApi.ArticleAddReq) (out interface{}, err error) {
	// 没有描述的时候，自动从文章内容获取描述
	if len(in.Description) == 0 {
		description, err := Article().getDescriptionByBody(ctx, in.Body)
		if err != nil {
			return nil, err
		}
		in.Description = description
	}
	// 没有缩略图时，自动获取文章的第一张图作为缩略图
	if len(in.Thumb) == 0 {
		thumb, err := Article().getThumbByBody(ctx, in.Body)
		if err != nil {
			return nil, err
		}
		in.Thumb = thumb
	}
	if len(in.Thumb) > 0 {
		//有缩略图，自动勾选in.FlagP
		in.FlagP = 1
	}
	// 构建flag
	flagStr, err := Article().buildFlagData(ctx, in.FlagP, in.FlagT, in.FlagR)
	if err != nil {
		return nil, err
	}
	// 关键词逗号替换成英文逗号
	keyword, err := Article().buildKeywordData(ctx, in.Keyword)
	if err != nil {
		return nil, err
	}
	id, err := dao.CmsArticle.Ctx(ctx).Data(g.Map{
		"title":       in.Title,
		"channelId":   in.ChannelId,
		"keyword":     keyword,
		"description": in.Description,
		"flag":        flagStr,
		"status":      in.Status,
		"thumb":       in.Thumb,
		"copyFrom":    in.CopyFrom,
		"clickNum":    in.ClickNum,
	}).InsertAndGetId()
	if err != nil {
		return nil, err
	}
	if id > 0 {
		_, err := dao.CmsArticleBody.Ctx(ctx).Data(g.Map{
			"articleId": id,
			"channelId": in.ChannelId,
			"body":      in.Body,
		}).Insert()
		if err != nil {
			return nil, err
		}
	}
	return
}

func (s *sArticle) Edit(ctx context.Context, in *backendApi.ArticleEditReq) (out interface{}, err error) {
	one, err := dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, in.Id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, gerror.New("文章不存在")
	}
	// 没有描述的时候，自动从文章内容获取描述
	if len(in.Description) == 0 {
		description, err := Article().getDescriptionByBody(ctx, in.Body)
		if err != nil {
			return nil, err
		}
		in.Description = description
	}
	// 没有缩略图时，自动获取文章的第一张图作为缩略图
	if len(in.Thumb) == 0 {
		thumb, err := Article().getThumbByBody(ctx, in.Body)
		if err != nil {
			return nil, err
		}
		in.Thumb = thumb
	}
	if len(in.Thumb) > 0 {
		//有缩略图，自动勾选in.FlagP
		in.FlagP = 1
	}
	// 构建flag
	flagStr, err := Article().buildFlagData(ctx, in.FlagP, in.FlagT, in.FlagR)
	if err != nil {
		return nil, err
	}
	// 关键词逗号替换成英文逗号
	keyword, err := Article().buildKeywordData(ctx, in.Keyword)
	if err != nil {
		return nil, err
	}
	_, err = dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, in.Id).Data(g.Map{
		"title":       in.Title,
		"channelId":   in.ChannelId,
		"keyword":     keyword,
		"description": in.Description,
		"flag":        flagStr,
		"status":      in.Status,
		"thumb":       in.Thumb,
		"copyFrom":    in.CopyFrom,
		"clickNum":    in.ClickNum,
	}).Update()
	if err != nil {
		return nil, err
	}
	_, err = dao.CmsArticleBody.Ctx(ctx).Where(dao.CmsArticleBody.Columns().ArticleId, in.Id).Data(g.Map{
		"channelId": in.ChannelId,
		"body":      in.Body,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sArticle) buildFlagData(ctx context.Context, flagP, flagT, flagR int) (flagData string, err error) {
	var data []string
	if flagP == 1 {
		data = append(data, "p")
	}
	if flagT == 1 {
		data = append(data, "t")
	}
	if flagR == 1 {
		data = append(data, "r")
	}
	flagData = gstr.Implode(",", data)
	return
}

func (s *sArticle) buildKeywordData(ctx context.Context, keyword string) (newKeyword string, err error) {
	str := "，"
	if gstr.Contains(keyword, str) {
		keyword = gstr.Replace(keyword, str, ",")
	}
	newKeyword = gstr.TrimRightStr(keyword, ",", -1)
	return
}

func (s *sArticle) getDescriptionByBody(ctx context.Context, body string) (description string, err error) {
	text := ghtml.StripTags(body)
	description = gstr.SubStrRune(text, 0, 255)
	return
}

func (s *sArticle) getThumbByBody(ctx context.Context, body string) (thumb string, err error) {
	// 自动获取文章缩略图
	autoArtPic := service.Util().GetSetting("auto_art_pic")
	// 功能未开启
	if autoArtPic != "1" {
		return
	}
	compileImg, err := regexp.Compile(`<[img|IMG].*?src=[\'|\"](.*?(?:[\.gif|\.jpg|\.png]))[\'|\"].*?[\/]?>`)
	if err != nil {
		return "", err
	}
	img := compileImg.FindString(body)
	if len(img) == 0 {
		return "", err
	}
	explodeImg := gstr.Explode(" ", img)
	if len(explodeImg) > 0 {
		for _, value := range explodeImg {
			// 找到src字符串
			if gstr.Contains(value, "src=") {
				thumb = gstr.ReplaceIByArray(value, []string{"src=", "", "\"", ""})
			}
		}
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
