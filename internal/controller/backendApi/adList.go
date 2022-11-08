package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	AdList = cAdList{}
)

type cAdList struct{}

// Index 广告列表
func (c *cAdList) Index(ctx context.Context, req *backendApi.AdListIndexReq) (res *backendApi.AdListIndexRes, err error) {
	var adList []*model.AdListItem
	m := dao.CmsAd.Ctx(ctx).As("ad")
	if req.ChannelId > 0 {
		m = m.Where("ad.channel_id", req.ChannelId)
	}
	err = m.LeftJoin(dao.CmsAdChannel.Table(), "ad_channel", "ad_channel.id=ad.Channel_id").
		Fields("ad.*", "ad_channel.channel_name").
		OrderAsc("ad.sort").
		OrderAsc("ad.id").
		Page(req.Page, req.Size).Scan(&adList)
	if err != nil {
		return nil, err
	}
	total, _ := m.Count()
	for key, item := range adList {
		if item.Status == 0 {
			adList[key].StatusDesc = "已停用"
		} else if item.StartTime == item.EndTime {
			adList[key].StatusDesc = "长启用"
			adList[key].StartTime = "永久"
			adList[key].EndTime = "永久"
		} else if item.StartTime <= gtime.Datetime() && gtime.Datetime() <= item.EndTime {
			adList[key].StatusDesc = "显示中"
		} else if item.StartTime > gtime.Datetime() {
			adList[key].StatusDesc = "待生效"
		} else if item.EndTime < gtime.Datetime() {
			adList[key].StatusDesc = "已过期"
		}
		if item.ImgUrl == "" {
			adList[key].ImgUrl = service.Util().ImageOrDefaultUrl(item.ImgUrl)
		}
	}
	res = &backendApi.AdListIndexRes{
		List:  adList,
		Total: total,
		Page:  req.Page,
		Size:  req.Size,
	}
	return
}

// Add 添加广告
func (c *cAdList) Add(ctx context.Context, req *backendApi.AdListAddReq) (res *backendApi.AdListAddRes, err error) {
	_, err = service.AdList().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "添加成功", nil)
	return
}

// Edit 编辑广告
func (c *cAdList) Edit(ctx context.Context, req *backendApi.AdListEditReq) (res *backendApi.AdListEditRes, err error) {
	_, err = service.AdList().Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "编辑成功", nil)
	return
}

// Delete 删除广告
func (c *cAdList) Delete(ctx context.Context, req *backendApi.AdListDeleteReq) (res *backendApi.AdListDeleteRes, err error) {
	_, err = service.AdList().Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "删除成功", nil)
	return
}
