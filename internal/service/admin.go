// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type IAdmin interface {
	LoginVerify(ctx context.Context, in model.AdminLoginInput) (admin *entity.CmsAdmin, err error)
	GetUserByUserNamePassword(ctx context.Context, in model.AdminLoginInput) g.Map
	GetRoleIdsByAccountId(accountId string) []gdb.Value
}

var localAdmin IAdmin

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
