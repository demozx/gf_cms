package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

// Login 管理员登录
func (c *cAdmin) Login(ctx context.Context, req *backend.AdminLoginReq) (res *backend.AdminLoginRes, err error) {
	//初始化系统管理员
	service.Admin().InitAdminUser(ctx)

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

// Add 添加管理员
func (c *cAdmin) Add(ctx context.Context, req *backend.AdminAddReq) (res *backend.AdminAddRes, err error) {
	var roleIdTitleArr []*model.RoleTitle
	err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Type, "backend").Where(dao.CmsRole.Columns().IsEnable, 1).Scan(&roleIdTitleArr)
	if err != nil {
		return nil, err
	}
	err = g.RequestFromCtx(ctx).Response.WriteTpl("backend/admin/add.html", g.Map{
		"roleIdTitleArr": roleIdTitleArr,
	})
	return
}

// Edit 编辑管理员
func (c *cAdmin) Edit(ctx context.Context, req *backend.AdminEditReq) (res *backend.AdminEditRes, err error) {
	var admin *entity.CmsAdmin
	err = dao.CmsAdmin.Ctx(ctx).Where(dao.CmsAdmin.Columns().Id, req.Id).Scan(&admin)
	if err != nil {
		return nil, err
	}
	if admin == nil {
		return nil, gerror.New("数据不存在")
	}
	var roleIdTitleArr []*model.RoleTitle
	err = dao.CmsRole.Ctx(ctx).Where(dao.CmsRole.Columns().Type, "backend").Where(dao.CmsRole.Columns().IsEnable, 1).Scan(&roleIdTitleArr)
	if err != nil {
		return nil, err
	}
	var roleAccounts []*entity.CmsRoleAccount
	err = dao.CmsRoleAccount.Ctx(ctx).Where(dao.CmsRoleAccount.Columns().AccountId, req.Id).Scan(&roleAccounts)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	err = g.RequestFromCtx(ctx).Response.WriteTpl("backend/admin/edit.html", g.Map{
		"roleIdTitleArr": roleIdTitleArr,
		"admin":          admin,
		"roleAccounts":   roleAccounts,
	})
	return
}
