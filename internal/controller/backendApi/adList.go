package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
)

var (
	AdList = cAdList{}
)

type cAdList struct{}

func (c *cAdList) Index(ctx context.Context, req *backendApi.AdChannelIndexReq) (res *backendApi.AdChannelIndexRes, err error) {
	return
}
