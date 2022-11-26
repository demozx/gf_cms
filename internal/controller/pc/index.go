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

func (c *cIndex) Index(ctx context.Context, req *pc.IndexReq) (res *pc.IndexRes, err error) {
	err = service.Response().View(ctx, "", g.Map{})
	if err != nil {
		return nil, err
	}
	return
}
