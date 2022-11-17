package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// 图集模型
type cImage struct{}

var (
	Image = cImage{}
)

func (c *cImage) Index(ctx context.Context, req *backendApi.ImageListReq) (res *backendApi.ImageListRes, err error) {
	var in *model.ImageGetListInPut
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	list, err := service.Image().BackendImageGetList(ctx, in)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "返回成功", list)
	return
}
