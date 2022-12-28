package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/consts"
)

var (
	RouterBeautify = cRouterBeautify{}
)

type cRouterBeautify struct{}

// About 路由美化-关于我们
func (c *cRouterBeautify) About(ctx context.Context, req *pc.RouterBeautifyReq) (res *pc.IndexRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: consts.AbortChannelTid})
	if err != nil {
		return nil, err
	}
	return
}
