package backend

import (
	"context"
	"gf_cms/api/backend"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Index(ctx context.Context, req *backend.WelcomeReq) (res *backend.WelcomeRes, err error) {
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("welcome/index.html", g.Map{})
	return
}
