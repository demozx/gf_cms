package image

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
)

type (
	sImage struct{}
)

var (
	insImage = sImage{}
)

func init() {
	service.RegisterImage(New())
}

func New() *sImage {
	return &sImage{}
}

func Image() *sImage {
	return &insImage
}

func (s *sImage) BackendImageGetList(ctx context.Context, in *model.ImageGetListInPut) (out *model.ImageGetListOutPut, err error) {
	m := dao.CmsImage.Ctx(ctx).As("image").OrderAsc("image.sort").OrderDesc("image.id")
	out = &model.ImageGetListOutPut{
		Page: in.Page,
		Size: in.Size,
	}
	if in.ChannelId > 0 {
		m = m.Where("image.channel_id", in.ChannelId)
	}
	if in.StartAt != "" && in.EndAt != "" {
		m = m.WhereGTE("image.created_at", in.StartAt).WhereLT("image.created_at", in.EndAt)
	}
	if in.Keyword != "" {
		m = m.WhereLike("image.title", "%"+in.Keyword+"%")
	}
	listModel := m.LeftJoin(dao.CmsChannel.Table(), "channel", "channel.id=image.channel_id").
		Fields("image.*, channel.name channel_name").
		Page(in.Page, in.Size)
	var list []*model.ImageListItem
	err = listModel.Scan(&list)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	if err = listModel.Scan(&out.List); err != nil {
		return out, err
	}
	for key, item := range out.List {
		thumb := service.Util().ImageOrDefaultUrl("")
		var otherImages []string
		if len(item.Images) > 0 {
			thumb = item.Images[0]
			otherImages = item.Images[1:]
		}
		out.List[key].Thumb = thumb             // 主图
		out.List[key].OtherImages = otherImages // 其他图
	}
	return
}
