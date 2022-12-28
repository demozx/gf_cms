package pc

import (
	"context"
	"gf_cms/api/pc"
)

var (
	RouterBeautify = cRouterBeautify{}
)

type cRouterBeautify struct{}

// About 路由美化-关于我们
func (c *cRouterBeautify) About(ctx context.Context, req *pc.RouterBeautifyAboutReq) (res *pc.IndexRes, err error) {
	_, err = SinglePage.Detail(ctx, &pc.SinglePageReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return
}
