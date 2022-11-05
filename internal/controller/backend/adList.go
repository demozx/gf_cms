package backend

import (
	"context"
	"gf_cms/api/backend"
)

var (
	AdList = cAdList{}
)

type cAdList struct{}

// Index 后台广告分类列表
func (c *cAdList) Index(ctx context.Context, req *backend.AdListIndexReq) (res *backend.AdListIndexRes, err error) {

	return
}
