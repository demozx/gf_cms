package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

// Login 管理员登录
func (c *cAdmin) Login(ctx context.Context, req *backend.AdminLoginReq) (res *backend.AdminLoginRes, err error) {
	var adminSession, _ = g.RequestFromCtx(ctx).Session.Get(consts.AdminSessionKeyPrefix)
	if !adminSession.IsEmpty() {
		// 有session，已经登录过
		var backendPrefix = util.Util().BackendPrefix()
		g.RequestFromCtx(ctx).Response.RedirectTo("/" + backendPrefix)
	}

	err = g.RequestFromCtx(ctx).Response.WriteTpl("backend/admin/login.html")
	if err != nil {
		panic(err)
	}
	return
}

// Index 管理员列表
func (c *cAdmin) Index(ctx context.Context, req *backend.AdminIndexReq) (res *backend.AdminIndexRes, err error) {
	list, err := service.Admin().BackendAdminGetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})

	err = g.RequestFromCtx(ctx).Response.WriteTpl("backend/admin/index.html", g.Map{
		"list":     list,
		"pageInfo": service.PageInfo().LayUiPageInfo(ctx, list.Total, list.Size),
	})

	if err != nil {
		return nil, err
	}

	return
}
