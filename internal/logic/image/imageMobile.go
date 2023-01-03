package image

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// MobileHomeRecommendGoodsList 移动首页推荐产品图集列表
func (s *sImage) MobileHomeRecommendGoodsList(ctx context.Context, belongChannelId int) (out []*model.ImageListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, belongChannelId, true)
	if err != nil {
		return nil, err
	}
	err = dao.CmsImage.Ctx(ctx).WhereIn(dao.CmsImage.Columns().ChannelId, arrAllIds).OrderRandom().Where(dao.CmsImage.Columns().Status, 1).Limit(4).Scan(&out)
	if err != nil {
		return nil, err
	}
	for key, item := range out {
		out[key], _ = service.Image().BuildThumb(ctx, item)
		out[key].Router, _ = service.GenUrl().DetailUrl(ctx, consts.ChannelModelImage, gconv.Int(item.Id))
	}
	return
}

// MobileHomeHonerList 移动首页荣誉资质列表
func (s *sImage) MobileHomeHonerList(ctx context.Context, channelId int) (out []*model.ImageListItem, err error) {
	err = dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().ChannelId, channelId).Where(dao.CmsImage.Columns().Status, 1).Limit(3).Scan(&out)
	if err != nil {
		return nil, err
	}
	for key, item := range out {
		out[key], _ = service.Image().BuildThumb(ctx, item)
		out[key].Router, _ = service.GenUrl().DetailUrl(ctx, consts.ChannelModelImage, gconv.Int(item.Id))
	}
	return
}
