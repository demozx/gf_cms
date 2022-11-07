package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	AdList = cAdList{}
)

type cAdList struct{}

// Index 后台广告分类列表
func (c *cAdList) Index(ctx context.Context, req *backend.AdListIndexReq) (res *backend.AdListIndexRes, err error) {
	var adChannel []*entity.CmsAdChannel
	err = dao.CmsAdChannel.Ctx(ctx).OrderAsc(dao.CmsAdChannel.Columns().Sort).OrderAsc(dao.CmsAdChannel.Columns().Id).Scan(&adChannel)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/ad/list/index.html", g.Map{
		"adChannel": adChannel,
	})
	if err != nil {
		return nil, err
	}
	return
}
