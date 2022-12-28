package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) Index(ctx context.Context, req *mobile.IndexReq) (res *mobile.IndexRes, err error) {
	err = service.Response().View(ctx, "/mobile/index/index.html", g.Map{})
	if err != nil {
		return nil, err
	}
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}
