package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Image = cImage{}
)

type cImage struct{}

func (c *cImage) Move(ctx context.Context, req *backend.ImageMoveReq) (res *backend.ImageMoveRes, err error) {
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelImage, 0)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/image/move.html", g.Map{
		"strIds":           req.StrIds,
		"channelModelTree": channelModelTree,
	})
	if err != nil {
		return nil, err
	}
	return
}
