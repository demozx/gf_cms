package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	SinglePage = cSinglePage{}
)

type cSinglePage struct{}

func (c *cSinglePage) Detail(ctx context.Context, req *mobile.SinglePageReq) (res *mobile.SinglePageRes, err error) {
	g.Dump("123")
	return
}
