package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/menu"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Index = cIndex{}
)

type cIndex struct{}

func (c *cIndex) Index(ctx context.Context, req *backend.IndexReq) (res *backend.IndexRes, err error) {
	var adminSession, _ = g.RequestFromCtx(ctx).Session.Get(consts.AdminSessionKeyPrefix)
	var cmsAdmin *entity.CmsAdmin
	err = adminSession.Scan(&cmsAdmin)
	if err != nil {
		panic(err)
	}
	accountId := gvar.New(cmsAdmin.Id).String()
	var backendMenu = menu.Menu().BackendMy(accountId)
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("backend/index/index.html", g.Map{
		"admin_session": gconv.Map(adminSession),
		"backend_menu":  backendMenu,
	})
	return
}
