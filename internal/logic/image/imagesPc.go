package image

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcHomeRecommendGoodsList pc首页推荐产品图集列表
func (s *sImage) PcHomeRecommendGoodsList(ctx context.Context, belongChannelId int) (out []*model.ImageListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, belongChannelId, true)
	if err != nil {
		return nil, err
	}
	err = dao.CmsImage.Ctx(ctx).WhereIn(dao.CmsImage.Columns().ChannelId, arrAllIds).OrderRandom().Where(dao.CmsImage.Columns().Status, 1).Scan(&out)
	if err != nil {
		return nil, err
	}
	for key, item := range out {
		out[key], _ = service.Image().BuildThumb(ctx, item)
		out[key].Router, _ = service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelImage, gconv.Int(item.Id))
	}
	return
}

func (s *sImage) PcHomeGoodsGroupList(ctx context.Context, belongChannelId int) (out [][]*model.ImageListItem, err error) {
	arrAllIds, err := service.Channel().GetChildIds(ctx, belongChannelId, false)
	var list = make([][]*model.ImageListItem, 0, len(arrAllIds))
	for _, channelId := range arrAllIds {
		var imageListItems []*model.ImageListItem
		err := dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().ChannelId, channelId).
			Where(dao.CmsImage.Columns().Status, 1).
			OrderAsc(dao.CmsImage.Columns().Sort).
			OrderDesc(dao.CmsImage.Columns().Id).
			Limit(3).
			Scan(&imageListItems)
		if err != nil {
			return nil, err
		}
		if len(imageListItems) > 0 {
			for key, item := range imageListItems {
				imageListItems[key].Router, _ = service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelImage, gconv.Int(item.Id))
				imageListItems[key], _ = service.Image().BuildThumb(ctx, item)
			}
		}
		list = append(list, imageListItems)
	}
	out = list
	return
}
