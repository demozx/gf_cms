package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/admin"
	"gf_cms/internal/logic/auth"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/do"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Admin = cAdmin{}
)

type cAdmin struct{}

// Login 管理员登录
func (c *cAdmin) Login(ctx context.Context, req *backendApi.AdminLoginReq) (res *backendApi.AdminLoginRes, err error) {
	adminInfo, err := admin.Admin().LoginVerify(ctx, model.AdminLoginInput{
		Username:   req.Username,
		Password:   req.Password,
		CaptchaStr: req.CaptchaStr,
		CaptchaId:  req.CaptchaId,
	})
	if err != nil {
		return nil, err
	}
	//如果登录的是系统用户
	if adminInfo.IsSystem == 1 {
		//给系统角色赋予全部的权限
		go func() {
			var systemRole *entity.CmsRole
			err = dao.CmsRole.Ctx(ctx).Where(do.CmsRole{IsSystem: 1}).Scan(&systemRole)
			if err != nil {
				return
			}
			//给系统角色增加权限
			_, err = dao.CmsRulePermissions.Ctx(ctx).Where(dao.CmsRulePermissions.Columns().V0, systemRole.Id).Delete()
			if err != nil {
				return
			}
			allViewPermissionsArray := service.Permission().GetAllViewPermissionsArray()
			allApiPermissionsArray := service.Permission().GetAllApiPermissionsArray()
			insertData := make([]interface{}, 0)
			for _, permission := range allViewPermissionsArray {
				row := g.Map{
					"p_type": "p",
					"v0":     systemRole.Id,
					"v1":     "backend",
					"v2":     permission,
				}
				insertData = append(insertData, row)
			}
			for _, permission := range allApiPermissionsArray {
				row := g.Map{
					"p_type": "p",
					"v0":     systemRole.Id,
					"v1":     "backend_api",
					"v2":     permission,
				}
				insertData = append(insertData, row)
			}
			_, err := dao.CmsRulePermissions.Ctx(ctx).Data(insertData).Insert()
			if err != nil {
				return
			}
		}()
	}
	res = &backendApi.AdminLoginRes{}
	//生成token
	res.Token, res.Expire = auth.Auth().JWTAuth().LoginHandler(ctx)
	//g.Dump(g.Map{
	//	"Token":  res.Token,
	//	"Expire": res.Expire,
	//})
	// 记录session：自己定义的，因为一般后台登录用session
	err = g.RequestFromCtx(ctx).Session.Set(consts.AdminSessionKeyPrefix, g.Map{
		"Token":    res.Token,
		"Id":       adminInfo.Id,
		"Username": adminInfo.Username,
		"Name":     adminInfo.Name,
	})
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.WriteJsonExit(g.Map{
		"code":    0,
		"message": "登录成功",
		"data":    res,
	})
	return
}

// Logout 管理员登出
func (c *cAdmin) Logout(ctx context.Context, req *backendApi.AdminLogoutReq) (res *backendApi.AdminLogoutRes, err error) {
	//清除session
	err = g.RequestFromCtx(ctx).Session.Remove(consts.AdminSessionKeyPrefix)
	if err != nil {
		return nil, err
	}
	//清除token
	auth.Auth().JWTAuth().LogoutHandler(ctx)
	return
}

// ClearCache 清除公共缓存
func (c *cAdmin) ClearCache(ctx context.Context, req *backendApi.AdminClearCacheReq) (res *backendApi.AdminClearCacheRes, err error) {
	_, err = util.Util().ClearPublicCache()
	if err != nil {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
			"code":    1,
			"message": err.Error(),
		})
	} else {
		service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "清理成功", g.Map{})
	}
	return
}

// Add 添加管理员
func (c *cAdmin) Add(ctx context.Context, req *backendApi.AdminAddReq) (res *backendApi.AdminAddRes, err error) {
	_, err = service.Admin().BackendApiAdminAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// Edit 添加管理员
func (c *cAdmin) Edit(ctx context.Context, req *backendApi.AdminEditReq) (res *backendApi.AdminEditRes, err error) {
	_, err = service.Admin().BackendApiAdminEdit(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}

// Status 修改状态
func (c *cAdmin) Status(ctx context.Context, req *backendApi.AdminStatusReq) (res *backendApi.AdminStatusRes, err error) {
	_, err = service.Admin().BackendApiAdminStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

// Delete 删除
func (c *cAdmin) Delete(ctx context.Context, req *backendApi.AdminDeleteReq) (res *backendApi.AdminDeleteRes, err error) {
	_, err = service.Admin().BackendApiAdminDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

// DeleteBatch 批量删除
func (c *cAdmin) DeleteBatch(ctx context.Context, req *backendApi.AdminDeleteBatchReq) (res *backendApi.AdminDeleteBatchRes, err error) {
	_, err = service.Admin().BackendApiAdminDeleteBatch(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}
