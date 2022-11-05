package backendApi

import (
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

var (
	AdChannel = cAdChannel{}
)

type cAdChannel struct{}

func (c *cAdChannel) Index(ctx context.Context, req *backendApi.AdChannelIndexReq) (res *backendApi.AdChannelIndexRes, err error) {
	var adChannelList []*model.AdChannelListItem
	m := dao.CmsAdChannel.Ctx(ctx).OrderAsc(dao.CmsAdChannel.Columns().Sort).OrderAsc(dao.CmsAdChannel.Columns().Id)
	err = m.Page(req.Page, req.Size).Scan(&adChannelList)
	if err != nil {
		return nil, err
	}
	count, err := m.Count()
	if err != nil {
		return nil, err
	}
	res = &backendApi.AdChannelIndexRes{
		Page:  req.Page,
		Size:  req.Size,
		List:  adChannelList,
		Total: count,
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "获取成功", res)
	return
}

func (c *cAdChannel) Add(ctx context.Context, req *backendApi.AdChannelAddReq) (res *backendApi.AdChannelAddRes, err error) {
	_, err = service.AdChannel().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "添加成功", g.Map{})
	return
}

func (c *cAdChannel) Edit(ctx context.Context, req *backendApi.AdChannelEditReq) (res *backendApi.AdChannelEditRes, err error) {
	_, err = service.AdChannel().Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "编辑成功", g.Map{})
	return
}

func (c *cAdChannel) Delete(ctx context.Context, req *backendApi.AdChannelDeleteReq) (res *backendApi.AdChannelDeleteRes, err error) {
	_, err = service.AdChannel().Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "删除成功", g.Map{})
	return
}

func (c *cAdChannel) Sort(ctx context.Context, req *backendApi.AdChannelSortReq) (res *backendApi.AdChannelSortRes, err error) {
	_, err = service.AdChannel().Sort(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "排序成功", g.Map{})
	return
}
