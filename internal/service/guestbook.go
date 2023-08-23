// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/api/backendApi"
	"gf_cms/api/mobileApi"
	"gf_cms/api/pcApi"
	"gf_cms/internal/model"
)

type (
	IGuestbook interface {
		// MobileSubmit mobile提交留言
		MobileSubmit(ctx context.Context, in *mobileApi.GuestbookReq) (out *mobileApi.GuestbookRes, err error)
		// PcSubmit pc提交留言
		PcSubmit(ctx context.Context, in *pcApi.GuestbookReq) (out *pcApi.GuestbookRes, err error)
		// GetAddressByIp 根据ip获取归属地
		GetAddressByIp(ctx context.Context, ip string) (address string, err error)
		// SendEmail 发送留言邮件
		SendEmail(ctx context.Context, guestbookId int64) (out interface{}, err error)
		// BackendGetList 留言板列表
		BackendGetList(ctx context.Context, in *backend.GuestbookIndexReq) (out *model.GuestbookGetListOutput, err error)
		// BackendApiStatus 修改留言状态
		BackendApiStatus(ctx context.Context, in *backendApi.GuestbookStatusReq) (out interface{}, err error)
		// BackendApiBatchDelete 批量删除留言
		BackendApiBatchDelete(ctx context.Context, in *backendApi.GuestbookDeleteReq) (out interface{}, err error)
	}
)

var (
	localGuestbook IGuestbook
)

func Guestbook() IGuestbook {
	if localGuestbook == nil {
		panic("implement not found for interface IGuestbook, forgot register?")
	}
	return localGuestbook
}

func RegisterGuestbook(i IGuestbook) {
	localGuestbook = i
}
