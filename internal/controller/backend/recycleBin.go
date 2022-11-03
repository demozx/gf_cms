package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	RecycleBin = cRecycleBin{}
)

type cRecycleBin struct{}

// Index 回收站列表
func (c *cRecycleBin) Index(ctx context.Context, req *backend.RecycleBinIndexReq) (res *backend.RecycleBinIndexRes, err error) {
	if req.Type == "" {
		req.Type = consts.ChannelModelArticle
	}
	if req.Type == consts.ChannelModelArticle {
		_, err = service.RecycleBin().ModelArticle(ctx, req)
	} else {
		return nil, gerror.New("不支持的模型")
	}
	if err != nil {
		return nil, err
	}
	return
}
