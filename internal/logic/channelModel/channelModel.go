package channelModel

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
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

func (*sChannelModel) ModelArticle(ctx context.Context, in *backend.ChannelModelIndexReq) (out []*model.ChannelBackendApiListItem, err error) {
	channelTree, err := service.Channel().BackendChannelTree(ctx, in.ChannelId)
	err = service.Response().View(ctx, "backend/channel_model/"+in.Type+".html", g.Map{
		"channelTree": channelTree,
		"modelType":   in.Type,
		"modelMap":    service.Channel().BackendModelMap(),
		"channel_id":  in.ChannelId,
		"withTab":     in.WithTab,
	})
	return
}
