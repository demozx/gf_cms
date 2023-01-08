package mobileApi

import (
	"context"
	"gf_cms/api/mobileApi"
	"gf_cms/internal/service"
)

var (
	Guestbook = cGuestbook{}
)

type cGuestbook struct{}

// Submit 提交留言
func (c *cGuestbook) Submit(ctx context.Context, req *mobileApi.GuestbookReq) (res *mobileApi.GuestbookRes, err error) {
	_, err = service.Guestbook().MobileSubmit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "提交成功")
	return
}
