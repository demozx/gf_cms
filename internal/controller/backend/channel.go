package backend

import (
	"context"
	"gf_cms/api/backend"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

func (c *cChannel) Index(ctx context.Context, req *backend.AdminLoginReq) (res *backend.AdminLoginRes, err error) {
	return
}
