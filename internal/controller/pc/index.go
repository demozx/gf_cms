package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

// Index pc首页
func (c *cIndex) Index(ctx context.Context, req *pc.IndexReq) (res *pc.IndexRes, err error) {
	// 导航栏
	navigation, err := service.Channel().PcNavigation(ctx)
	if err != nil {
		return nil, err
	}
	g.Dump("navigation", navigation)
	// banner广告

	err = service.Response().View(ctx, "/pc/index/index.html", g.Map{
		"navigation": navigation,
	})
	if err != nil {
		return nil, err
	}
	return
}
