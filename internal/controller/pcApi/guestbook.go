package pcApi

import (
	"context"
	"gf_cms/api/pcApi"
	"gf_cms/internal/service"
)

var (
	Guestbook = cGuestbook{}
)

type cGuestbook struct{}

// Submit 提交留言
func (c *cGuestbook) Submit(ctx context.Context, req *pcApi.GuestbookReq) (res *pcApi.GuestbookRes, err error) {
	_, err = service.Guestbook().PcSubmit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "提交成功")
	return
}
