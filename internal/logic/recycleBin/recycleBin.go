package recycleBin

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

type sRecycleBin struct{}

var (
	insRecycleBin = sRecycleBin{}
)

func init() {
	service.RegisterRecycleBin(New())
}

func New() *sRecycleBin {
	return &sRecycleBin{}
}

// RecycleBin 回收站
func RecycleBin() *sRecycleBin {
	return &insRecycleBin
}

func (*sRecycleBin) ModelArticle(ctx context.Context, req *backend.RecycleBinIndexReq) (out []*model.ChannelBackendApiListItem, err error) {
	err = service.Response().View(ctx, "backend/recycle_bin/article/index.html", g.Map{
		"modelMap": service.Channel().BackendModelMap(),
	})
	return
}
