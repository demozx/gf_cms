// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model/entity"
)

type (
	IShortcut interface {
		// BackendIndex 列表
		BackendIndex(ctx context.Context) (out []*entity.CmsShortcut, err error)
		// BackendEdit 后台编辑返回信息
		BackendEdit(ctx context.Context, in *backend.ShortcutEditReq) (out *entity.CmsShortcut, err error)
		// BackendApiAdd 添加快捷方式
		BackendApiAdd(ctx context.Context, in *backendApi.ShortcutAddReq) (out interface{}, err error)
		// BackendApiEdit 编辑快捷方式
		BackendApiEdit(ctx context.Context, in *backendApi.ShortcutEditReq) (out interface{}, err error)
		// BackendApiBatchDelete 删除快捷方式
		BackendApiBatchDelete(ctx context.Context, in *backendApi.ShortcutBatchDeleteReq) (out interface{}, err error)
		// BackendApiSort 排序
		BackendApiSort(ctx context.Context, in *backendApi.ShortcutSortReq) (out interface{}, err error)
	}
)

var (
	localShortcut IShortcut
)

func Shortcut() IShortcut {
	if localShortcut == nil {
		panic("implement not found for interface IShortcut, forgot register?")
	}
	return localShortcut
}

func RegisterShortcut(i IShortcut) {
	localShortcut = i
}
