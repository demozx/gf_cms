package channelModel

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sChannelModel struct{}

var (
	insChannelModel = sChannelModel{}
)

func init() {
	service.RegisterChannelModel(New())
}

func New() *sChannelModel {
	return &sChannelModel{}
}

func ChannelModel() *sChannelModel {
	return &insChannelModel
}

// ModelArticle 文章模型列表
func (s *sChannelModel) ModelArticle(ctx context.Context, in *backend.ChannelModelIndexReq) (out []*model.ChannelBackendApiListItem, err error) {
	_, err = ChannelModel().checkChannel(ctx, in.ChannelId)
	if err != nil {
		return nil, err
	}
	channelTree, err := service.Channel().BackendChannelModelTree(ctx, in.Type, in.ChannelId)
	if err != nil {
		return nil, err
	}
	recycleBin, err := service.Util().GetSetting("recycle_bin")
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/article/index.html", g.Map{
		"channelTree": channelTree,
		"modelType":   in.Type,
		"modelMap":    service.Channel().BackendModelCanAddMap(),
		"channelId":   in.ChannelId,
		"withTab":     in.WithTab,
		"deleteType":  recycleBin,
	})
	if err != nil {
		return nil, err
	}
	return
}

// ModelImage 图集模型列表
func (s *sChannelModel) ModelImage(ctx context.Context, in *backend.ChannelModelIndexReq) (out []*model.ChannelBackendApiListItem, err error) {
	_, err = ChannelModel().checkChannel(ctx, in.ChannelId)
	if err != nil {
		return nil, err
	}
	channelTree, err := service.Channel().BackendChannelModelTree(ctx, in.Type, in.ChannelId)
	err = service.Response().View(ctx, "backend/channel_model/image/index.html", g.Map{
		"channelTree": channelTree,
		"modelType":   in.Type,
		"modelMap":    service.Channel().BackendModelCanAddMap(),
		"channelId":   in.ChannelId,
		"withTab":     in.WithTab,
	})
	return
}

// 检测频道id是否存在
func (s *sChannelModel) checkChannel(ctx context.Context, channelId int) (out interface{}, err error) {
	if channelId > 0 {
		one, err := dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).One()
		if err != nil {
			return nil, err
		}
		if one == nil {
			return nil, gerror.New("频道不存在")
		}
	}
	return
}

// GetDetailOneByChannelId 根据栏目id，详情页id，获取一条对应的详情
func (s *sChannelModel) GetDetailOneByChannelId(ctx context.Context, channelId uint, detailId int64) (out interface{}, err error) {
	var channelInfo *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channelInfo)
	if err != nil {
		return nil, err
	}
	if channelInfo == nil {
		return nil, gerror.New("栏目不存在")
	}
	switch channelInfo.Model {
	case consts.ChannelModelArticle:
		var article *entity.CmsArticle
		err := dao.CmsArticle.Ctx(ctx).Where(dao.CmsArticle.Columns().Id, detailId).Scan(&article)
		if err != nil {
			return nil, err
		}
		if article == nil {
			return nil, gerror.New("文章不存在 " + gconv.String(detailId))
		}
		out = article
	case consts.ChannelModelImage:
		var image *entity.CmsImage
		err := dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, detailId).Scan(&image)
		if err != nil {
			return nil, err
		}
		if image == nil {
			return nil, gerror.New("图集不存在 " + gconv.String(detailId))
		}
		out = image
	case consts.ChannelModelSinglePage:
		return nil, gerror.New("单页模型不支持详情内容")
	}
	return
}
