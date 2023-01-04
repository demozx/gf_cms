package mobile

import (
	"context"
	"gf_cms/api/mobile"
	"gf_cms/internal/consts"
)

var (
	RouterBeautify = cRouterBeautify{}
)

type cRouterBeautify struct{}

// About 路由美化-关于我们
func (c *cRouterBeautify) About(ctx context.Context, req *mobile.RouterBeautifyReq) (res *mobile.RouterBeautifyRes, err error) {
	_, err = SinglePage.Detail(ctx, &mobile.SinglePageReq{Id: consts.AboutChannelId})
	if err != nil {
		return nil, err
	}
	return
}
