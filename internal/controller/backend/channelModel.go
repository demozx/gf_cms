package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ChannelModel = cChannelModel{}
)

type cChannelModel struct{}

// Index 栏目分类列表
func (c *cChannelModel) Index(ctx context.Context, req *backend.ChannelModelIndexReq) (res *backend.ChannelModelIndexRes, err error) {
	if req.Type == "" {
		req.Type = consts.ChannelModelArticle
	}
	if req.Type == consts.ChannelModelArticle {
		_, err = service.ChannelModel().ModelArticle(ctx, req)
	} else {
		return nil, gerror.New("不支持的模型")
	}
	if err != nil {
		return nil, err
	}
	return
}
