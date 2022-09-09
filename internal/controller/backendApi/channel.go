package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

// Index 获取后台栏目分类接口数据
func (c *cChannel) Index(ctx context.Context, req *backendApi.ChannelIndexApiReq) (res *backendApi.ChannelIndexApiRes, err error) {
	data, err := service.Channel().Index(ctx)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, 0, "获取成功", data)
	return
}
