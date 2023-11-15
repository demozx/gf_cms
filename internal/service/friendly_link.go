// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model/entity"
)

type (
	IFriendlyLink interface {
		// BackendApiStatus 修改友情链接状态
		BackendApiStatus(ctx context.Context, req *backendApi.FriendlyLinkStatusReq) (res interface{}, err error)
		// BackendApiAdd 添加友情链接
		BackendApiAdd(ctx context.Context, req *backendApi.FriendlyLinkAddReq) (res interface{}, err error)
		// BackendApiEdit 编辑友情链接
		BackendApiEdit(ctx context.Context, req *backendApi.FriendlyLinkEditReq) (res interface{}, err error)
		// BackendApiSort 友情链接排序
		BackendApiSort(ctx context.Context, req *backendApi.FriendlyLinkSortReq) (res interface{}, err error)
		// BackendApiBatchDelete 批量删除
		BackendApiBatchDelete(ctx context.Context, req *backendApi.FriendlyLinkBatchDeleteReq) (res interface{}, err error)
		// PcList pc友情链接
		PcList(ctx context.Context) (out []*entity.CmsFriendlyLink, err error)
	}
)

var (
	localFriendlyLink IFriendlyLink
)

func FriendlyLink() IFriendlyLink {
	if localFriendlyLink == nil {
		panic("implement not found for interface IFriendlyLink, forgot register?")
	}
	return localFriendlyLink
}

func RegisterFriendlyLink(i IFriendlyLink) {
	localFriendlyLink = i
}
