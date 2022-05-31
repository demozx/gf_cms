package backend

import (
	"context"
	"gf_cms/api/admin"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

func (c *cChannel) Index(ctx context.Context, req *admin.LoginReq) (res *admin.LoginRes, err error) {
	return
}
