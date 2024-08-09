// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IPermission interface {
		// BackendAll 获取后台全部权限（view和api）
		BackendAll() (res []model.PermissionAllItem, err error)
		// BackendViewAll Backend 获取全部后台权限
		BackendViewAll() (res []model.PermissionGroups, err error)
		// BackendApiAll Backend 获取全部后台接口权限
		BackendApiAll() (res []model.PermissionGroups, err error)
		// BackendMyView 获取我的所有后台视图权限
		BackendMyView(accountId string) (myPermissions []gdb.Value, err error)
		// BackendMyApi 获取我的所有后台接口权限
		BackendMyApi(accountId string) (myPermissions []gdb.Value, err error)
		// GetAllViewPermissionsArray 获取全部视图权限数组
		GetAllViewPermissionsArray() (res []string, err error)
		// GetAllApiPermissionsArray 获取全部接口权限数组
		GetAllApiPermissionsArray() (res []string, err error)
		// BackendUserViewCan 检测后台用户有无视图的操作权限
		BackendUserViewCan(ctx context.Context, routePermission string) bool
		// BackendUserApiCan 检测后台用户有无接口的操作权限
		BackendUserApiCan(ctx context.Context, routePermission string) bool
		// BackendGetUserFromSession 从session获取当前登录用户
		BackendGetUserFromSession(ctx context.Context) (out *entity.CmsAdmin, err error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
