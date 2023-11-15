// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type (
	IAdmin interface {
		// LoginVerify 登录验证
		LoginVerify(ctx context.Context, in model.AdminLoginInput) (admin *entity.CmsAdmin, err error)
		GetUserByUserNamePassword(ctx context.Context, in model.AdminLoginInput) g.Map
		// GetRoleIdsByAccountId 获取用户的所有角色id
		GetRoleIdsByAccountId(accountId string) []gdb.Value
		// BackendAdminGetList 后台获取管理员列表
		BackendAdminGetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
		// BackendApiAdminAdd 添加管理员
		BackendApiAdminAdd(ctx context.Context, in *backendApi.AdminAddReq) (out interface{}, err error)
		// BackendApiAdminEdit 编辑
		BackendApiAdminEdit(ctx context.Context, in *backendApi.AdminEditReq) (out interface{}, err error)
		// BackendApiAdminStatus 修改自动状态
		BackendApiAdminStatus(ctx context.Context, in *backendApi.AdminStatusReq) (out interface{}, err error)
		// BackendApiAdminDelete 删除
		BackendApiAdminDelete(ctx context.Context, in *backendApi.AdminDeleteReq) (out interface{}, err error)
		// BackendApiAdminDeleteBatch 批量删除
		BackendApiAdminDeleteBatch(ctx context.Context, in *backendApi.AdminDeleteBatchReq) (out interface{}, err error)
		// InitAdminUser 初始化系统管理员
		InitAdminUser(ctx context.Context)
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
