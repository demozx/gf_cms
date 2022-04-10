package admin

import (
	"context"
	"gf_cms/api/admin"
	"gf_cms/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) Index(ctx context.Context, req *admin.LoginReq) (res *admin.LoginRes, err error) {
	var adminSession, _ = g.RequestFromCtx(ctx).Session.Get(consts.AdminSessionKeyPrefix)

	g.RequestFromCtx(ctx).Response.WriteTpl("index/index.html", g.Map{
		"admin_session": gconv.Map(adminSession),
	})
	return
}
