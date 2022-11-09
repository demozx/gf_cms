package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

var (
	Guestbook = cGuestbook{}
)

type cGuestbook struct{}

// Status 修改留言状态
func (c *cGuestbook) Status(ctx context.Context, req *backendApi.GuestbookStatusReq) (res *backendApi.GuestbookStatusRes, err error) {
	_, err = service.Guestbook().BackendApiStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

// BatchDelete 删除留言
func (c *cGuestbook) BatchDelete(ctx context.Context, req *backendApi.GuestbookDeleteReq) (res *backendApi.GuestbookDeleteRes, err error) {
	_, err = service.Guestbook().BackendApiBatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "删除成功")
	return
}
