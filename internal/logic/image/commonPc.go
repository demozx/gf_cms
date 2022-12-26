package image

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcPrevImage 上一篇图集
func (s *sImage) PcPrevImage(ctx context.Context, channelId int, imageId uint64) (out *model.ImageLink, err error) {
	var prevImage *entity.CmsImage
	err = dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().ChannelId, channelId).Where(dao.CmsImage.Columns().Status, 1).WhereLT(dao.CmsImage.Columns().Id, imageId).OrderAsc(dao.CmsImage.Columns().Sort).OrderDesc(dao.CmsImage.Columns().Id).Scan(&prevImage)
	if err != nil {
		return
	}
	if prevImage == nil {
		out = &model.ImageLink{Title: "无", Router: "javascript:;"}
	} else {
		url, err := service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelImage, gconv.Int(prevImage.Id))
		if err != nil {
			return nil, err
		}
		out = &model.ImageLink{Title: prevImage.Title, Router: url}
	}
	return
}

// PcNextImage 下一篇图集
func (s *sImage) PcNextImage(ctx context.Context, channelId int, imageId uint64) (out *model.ImageLink, err error) {
	var nextImage *entity.CmsImage
	err = dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().ChannelId, channelId).Where(dao.CmsImage.Columns().Status, 1).WhereGT(dao.CmsImage.Columns().Id, imageId).OrderAsc(dao.CmsImage.Columns().Sort).OrderAsc(dao.CmsImage.Columns().Id).Scan(&nextImage)
	if err != nil {
		return
	}
	if nextImage == nil {
		out = &model.ImageLink{Title: "无", Router: "javascript:;"}
	} else {
		url, err := service.GenUrl().PcDetailUrl(ctx, consts.ChannelModelImage, gconv.Int(nextImage.Id))
		if err != nil {
			return nil, err
		}
		out = &model.ImageLink{Title: nextImage.Title, Router: url}
	}
	return
}
