package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

type cRole struct{}

var Role = cRole{}

// Status 修改状态
func (c *cRole) Status(ctx context.Context, req *backendApi.RoleStatusReq) (res *backendApi.AdminStatusRes, err error) {
	_, err = service.Role().BackendApiRoleStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

// Delete 角色删除
func (c *cRole) Delete(ctx context.Context, req *backendApi.RoleDeleteReq) (res *backendApi.RoleDeleteRes, err error) {
	_, err = service.Role().BackendApiRoleDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cRole) DeleteBatch(ctx context.Context, req *backendApi.RoleDeleteBatchReq) (res *backendApi.RoleDeleteBatchRes, err error) {
	_, err = service.Role().BackendApiRoleDeleteBatch(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}
