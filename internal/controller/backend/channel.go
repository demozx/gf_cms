package backend

import (
	"context"
	"gf_cms/api/backend"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

func (c *cChannel) Index(ctx context.Context, req *backend.LoginReq) (res *backend.LoginRes, err error) {
	return
}
