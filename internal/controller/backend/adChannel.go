package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/dao"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	AdChannel = cAdChannel{}
)

type cAdChannel struct{}

// Index 后台广告分类列表
func (c *cAdChannel) Index(ctx context.Context, req *backend.AdChannelIndexReq) (res *backend.AdChannelIndexRes, err error) {
	err = service.Response().View(ctx, "/backend/ad/channel/index.html", g.Map{})
	return
}

func (c *cAdChannel) Edit(ctx context.Context, req *backend.AdChannelEditReq) (res *backend.AdChannelEditRes, err error) {
	var adChannel *entity.CmsAdChannel
	err = dao.CmsAdChannel.Ctx(ctx).Where(dao.CmsAdChannel.Columns().Id, req.Id).Scan(&adChannel)
	if err != nil {
		return nil, err
	}
	if adChannel == nil {
		return nil, gerror.New("广告分类不存在")
	}
	err = service.Response().View(ctx, "/backend/ad/channel/edit.html", g.Map{
		"adChannel": adChannel,
	})
	return
}
