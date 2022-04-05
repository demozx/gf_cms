package admin

import (
	"context"
	"gf_cms/api/admin"
	"gf_cms/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

func (c *cAdmin) Login(ctx context.Context, req *admin.LoginReq) (res *admin.LoginRes, err error) {
	var adminSession, _ = g.RequestFromCtx(ctx).Session.Get(consts.AdminSessionKeyPrefix)
	if !adminSession.IsEmpty() {
		// 有session，已经登录过
		var AdminPrefix, _ = g.Cfg().Get(ctx, "server.adminPrefix", "admin")
		g.RequestFromCtx(ctx).Response.RedirectTo("/" + AdminPrefix.String())
	}

	g.RequestFromCtx(ctx).Response.WriteTpl("admin/login.html")
	return
}
