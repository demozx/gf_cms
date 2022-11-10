package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

var (
	FriendlyLink = cFriendlyLink{}
)

type cFriendlyLink struct{}

// Status 修改友情链接状态
func (c *cFriendlyLink) Status(ctx context.Context, req *backendApi.FriendlyLinkStatusReq) (res *backendApi.FriendlyLinkStatusRes, err error) {
	_, err = service.FriendlyLink().BackendApiStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}
