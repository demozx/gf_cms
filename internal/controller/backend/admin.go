package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/util"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

func (c *cAdmin) Login(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
	var adminSession, _ = g.RequestFromCtx(ctx).Session.Get(consts.AdminSessionKeyPrefix)
	if !adminSession.IsEmpty() {
		// 有session，已经登录过
		var backendPrefix = util.Util().BackendPrefix()
		g.RequestFromCtx(ctx).Response.RedirectTo("/" + backendPrefix)
	}

	g.RequestFromCtx(ctx).Response.WriteTpl("backend/admin/login.html")
	return
}
