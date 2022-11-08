package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	AdList = cAdList{}
)

type cAdList struct{}

// Index 后台广告分类列表
func (c *cAdList) Index(ctx context.Context, req *backend.AdListIndexReq) (res *backend.AdListIndexRes, err error) {
	adChannel, err := service.AdChannel().GetAdChannelMap(ctx)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/ad/list/index.html", g.Map{
		"adChannel": adChannel,
	})
	if err != nil {
		return nil, err
	}
	return
}

// Add 添加广告
func (c *cAdList) Add(ctx context.Context, req *backend.AdListAddReq) (res *backend.AdListAddRes, err error) {
	adChannel, err := service.AdChannel().GetAdChannelMap(ctx)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/ad/list/add.html", g.Map{
		"adChannel": adChannel,
	})
	if err != nil {
		return nil, err
	}
	return
}

// Edit 编辑广告
func (c *cAdList) Edit(ctx context.Context, req *backend.AdListEditReq) (res *backend.AdListEditRes, err error) {
	adChannel, err := service.AdChannel().GetAdChannelMap(ctx)
	if err != nil {
		return nil, err
	}
	adInfo, err := service.AdList().GetAdInfoById(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/ad/list/edit.html", g.Map{
		"adChannel": adChannel,
		"adInfo":    adInfo,
	})
	if err != nil {
		return nil, err
	}
	return
}
