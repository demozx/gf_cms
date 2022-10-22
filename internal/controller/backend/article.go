package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

func (c *cArticle) Move(ctx context.Context, req *backend.ArticleMoveReq) (res *backend.ArticleMoveRes, err error) {
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelArticle, 0)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/article/move.html", g.Map{
		"strIds":           req.StrIds,
		"channelModelTree": channelModelTree,
	})
	if err != nil {
		return nil, err
	}
	return
}