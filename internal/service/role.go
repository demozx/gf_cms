// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
)

type (
	IRole interface {
		// BackendRoleGetList 获取角色列表
		BackendRoleGetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
		// BackendRoleGetOne 获取单个角色
		BackendRoleGetOne(ctx context.Context, in *backend.RoleEditReq) (out *model.RoleItem, err error)
		// BackendApiRoleStatus 修改角色状态
		BackendApiRoleStatus(ctx context.Context, in *backendApi.RoleStatusReq) (out interface{}, err error)
		// BackendApiRoleDelete 角色删除
		BackendApiRoleDelete(ctx context.Context, in *backendApi.RoleDeleteReq) (out interface{}, err error)
		// BackendApiRoleDeleteBatch 角色批量删除
		BackendApiRoleDeleteBatch(ctx context.Context, in *backendApi.RoleDeleteBatchReq) (out interface{}, err error)
		// BackendApiRoleAdd 添加角色
		BackendApiRoleAdd(ctx context.Context, in *backendApi.RoleAddReq) (out interface{}, err error)
		// BackendApiRoleEdit 编辑角色
		BackendApiRoleEdit(ctx context.Context, in *backendApi.RoleEditReq) (out interface{}, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
