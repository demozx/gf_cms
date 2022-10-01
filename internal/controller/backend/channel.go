package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

// Index 栏目分类列表
func (c *cChannel) Index(ctx context.Context, req *backend.ChannelIndexReq) (res *backend.ChannelIndexRes, err error) {
	err = service.Response().View(ctx, "backend/channel/index.html", g.Map{})
	if err != nil {
		return nil, err
	}
	return
}

// Add 添加
func (c *cChannel) Add(ctx context.Context, req *backend.ChannelAddReq) (res *backend.ChannelAddRes, err error) {
	channelTree, err := service.Channel().BackendChannelTree(ctx, req.Id)
	modelMap := service.Channel().BackendModelMap()
	err = service.Response().View(ctx, "backend/channel/add.html", g.Map{
		"channelTree": channelTree,
		"modelMap":    modelMap,
	})
	if err != nil {
		return nil, err
	}
	return
}

// Edit 编辑
func (c *cChannel) Edit(ctx context.Context, req *backend.ChannelEditReq) (res *backend.ChannelEditRes, err error) {
	channel, err := service.Channel().GetOneById(ctx, req.Id)
	channelTree, err := service.Channel().BackendChannelTree(ctx, channel.Pid)
	modelMap := service.Channel().BackendModelMap()
	err = service.Response().View(ctx, "backend/channel/edit.html", g.Map{
		"channelTree": channelTree,
		"modelMap":    modelMap,
		"channel":     channel,
	})
	if err != nil {
		return nil, err
	}
	return
}
