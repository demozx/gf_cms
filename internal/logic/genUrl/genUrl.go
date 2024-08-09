package genUrl

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	insGenUrl = sGenUrl{}
)

type sGenUrl struct{}

func init() {
	service.RegisterGenUrl(New())
}

func New() *sGenUrl {
	return &sGenUrl{}
}

func GenUrl() *sGenUrl {
	return &insGenUrl
}

// ChannelUrl 生成pc栏目url
// channelId 栏目id
// router 可传空，非空可减少一次查询
func (s *sGenUrl) ChannelUrl(ctx context.Context, channelId int, router string) (newRouter string, err error) {
	if router != "" {
		// 路由中有{id}字符串，替换成指定的id
		if gstr.Contains(router, "{id}") {
			newRouter = gstr.Replace(router, "{id}", gconv.String(channelId), 1)
		} else {
			newRouter = router
		}
	} else {
		var channel *model.ChannelNavigationListItem
		cacheKey := util.PublicCachePreFix + ":channel_url:" + gconv.String(channelId)
		exists, err := g.Redis().Do(ctx, "EXISTS", cacheKey)
		if err != nil {
			return "", err
		}
		if exists.Bool() {
			cached, err := g.Redis().Do(ctx, "GET", cacheKey)
			if err != nil {
				return "", err
			}
			if cached.String() == "" {
				return "", gerror.New("栏目不存在")
			}
			return cached.String(), nil
		} else {
			err := dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channel)
			if err != nil {
				return "", err
			}
			if channel == nil {
				return "", gerror.New("栏目不存在")
			}
		}
		// 根据频道类型处理url
		switch channel.Type {
		case 1:
			// 频道类型
			fallthrough
		case 2:
			// 单页类型
			newRouter = channel.ListRouter
			if gstr.Contains(channel.ListRouter, "{id}") {
				// 如果路由中有{id}，替换id
				newRouter, _ = service.GenUrl().ChannelUrl(ctx, gconv.Int(channel.Id), channel.ListRouter)
			}
		case 3:
			// 链接类型
			newRouter = channel.LinkUrl
		default:
			return "", gerror.New("栏目类型错误")
		}
		_, err = g.Redis().Do(ctx, "SET", cacheKey, newRouter)
		if err != nil {
			return "", err
		}
	}
	return
}

// DetailUrl 生成pc详情页url
func (s *sGenUrl) DetailUrl(ctx context.Context, model string, detailId int) (newRouter string, err error) {
	cacheKey := util.PublicCachePreFix + ":detail_url:" + model + ":" + gconv.String(detailId)
	cached, err := g.Redis().Do(ctx, "GET", cacheKey)
	if err != nil {
		return "", err
	}
	if !cached.IsEmpty() {
		return cached.String(), nil
	}
	channelId := 0
	if model == consts.ChannelModelArticle {
		var article *entity.CmsArticle
		err := dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, detailId).Scan(&article)
		if err != nil {
			return "", err
		}
		if article != nil {
			channelId = article.ChannelId
		}
	} else if model == consts.ChannelModelImage {
		var image *entity.CmsImage
		err := dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, detailId).Scan(&image)
		if err != nil {
			return "", err
		}
		if image != nil {
			channelId = image.ChannelId
		}
	} else {
		return "", nil
	}

	var channel *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Model, model).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channel)
	if err != nil {
		return "", err
	}
	if channel == nil || channel.Model == "" {
		return "", gerror.New("栏目或模型不存在")
	}
	// 路由中有{id}字符串，替换成指定的id
	if gstr.Contains(channel.DetailRouter, "{id}") {
		newRouter = gstr.Replace(channel.DetailRouter, "{id}", gconv.String(detailId), 1)
	} else {
		newRouter = channel.DetailRouter
	}
	_, err = g.Redis().Do(ctx, "SET", cacheKey, newRouter)
	if err != nil {
		return "", err
	}
	return
}
