package image

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"sync"
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
		allIds, err := service.Channel().GetChildIds(ctx, in.ChannelId, true)
		if err != nil {
			return nil, err
		}
		m = m.WhereIn("image.channel_id", allIds)
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

func (s *sImage) Sort(ctx context.Context, in []*model.ImageSortMap) (out interface{}, err error) {
	update, err := dao.CmsImage.Ctx(ctx).Save(in)
	if err != nil {
		return false, err
	}
	return update, nil
}

func (s *sImage) Flag(ctx context.Context, ids []int, flagType string) (out interface{}, err error) {
	if len(ids) == 1 {
		_, err = Image().singleFlag(ctx, ids[0], flagType, "auto")
		if err != nil {
			return nil, err
		}
	} else {
		var wg sync.WaitGroup
		for _, id := range ids {
			wg.Add(1)
			go func(id int) {
				_, err = Image().singleFlag(ctx, id, flagType, "open")
				wg.Done()
				if err != nil {
					return
				}
			}(id)
		}
		wg.Wait()
	}
	return
}

func (s *sImage) Status(ctx context.Context, ids []int) (out interface{}, err error) {
	if len(ids) == 1 {
		_, err = Image().singleStatus(ctx, ids[0], "auto")
		if err != nil {
			return nil, err
		}
	} else {
		var wg sync.WaitGroup
		for _, id := range ids {
			wg.Add(1)
			go func(id int) {
				_, err = Image().singleStatus(ctx, id, "open")
				wg.Done()
				if err != nil {
					return
				}
			}(id)
		}
		wg.Wait()
	}
	return
}

func (s *sImage) Delete(ctx context.Context, ids []int) (out interface{}, err error) {
	m := dao.CmsImage.Ctx(ctx).WhereIn(dao.CmsImage.Columns().Id, ids)
	if service.Util().GetSetting("recycle_bin") == "1" {
		_, err = m.Delete()
	} else {
		_, err = m.Unscoped().Delete()
	}

	if err != nil {
		return nil, err
	}
	return
}

func (s *sImage) Move(ctx context.Context, channelId int, ids []string) (out interface{}, err error) {
	if channelId <= 0 {
		return nil, gerror.New("频道ID错误")
	}
	if len(ids) == 0 {
		return nil, gerror.New("要移动的图集不能为空")
	}
	_, err = dao.CmsImage.Ctx(ctx).WhereIn(dao.CmsImage.Columns().Id, ids).Data(g.Map{
		dao.CmsImage.Columns().ChannelId: channelId,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sImage) Add(ctx context.Context, in *backendApi.ImageAddReq) (out interface{}, err error) {
	// 生成图片数组
	imagesArr, err := Image().buildImagesArr(ctx, in.Images)
	if err != nil {
		return nil, err
	}
	// 构建flag
	flagStr, err := Image().buildFlagData(ctx, in.FlagT, in.FlagR)
	if err != nil {
		return nil, err
	}
	_, err = dao.CmsImage.Ctx(ctx).Data(g.Map{
		"channelId":   in.ChannelId,
		"title":       in.Title,
		"images":      imagesArr,
		"description": in.Description,
		"flag":        flagStr,
		"status":      in.Status,
		"clickNum":    in.ClickNum,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sImage) Edit(ctx context.Context, in *backendApi.ImageEditReq) (out interface{}, err error) {
	// 生成图片数组
	imagesArr, err := Image().buildImagesArr(ctx, in.Images)
	if err != nil {
		return nil, err
	}
	// 构建flag
	flagStr, err := Image().buildFlagData(ctx, in.FlagT, in.FlagR)
	if err != nil {
		return nil, err
	}
	affected, err := dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, in.Id).Data(g.Map{
		"channelId":   in.ChannelId,
		"title":       in.Title,
		"images":      imagesArr,
		"description": in.Description,
		"flag":        flagStr,
		"status":      in.Status,
		"clickNum":    in.ClickNum,
	}).UpdateAndGetAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, gerror.New("图集不存在")
	}
	return
}

// BackendRecycleBinImageGetList 回收站图集列表
func (s *sImage) BackendRecycleBinImageGetList(ctx context.Context, in *model.ImageGetListInPut) (out *model.ImageGetListOutPut, err error) {
	m := dao.CmsImage.Ctx(ctx).
		As("image").OrderAsc("image.sort").OrderDesc("image.id").WhereNotNull("image.deleted_at").Unscoped()
	out = &model.ImageGetListOutPut{
		Page: in.Page,
		Size: in.Size,
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
	//g.Dump("out.List", out.List)
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

// BackendRecycleBinImageBatchDestroy 回收站-图集批量永久删除
func (s *sImage) BackendRecycleBinImageBatchDestroy(ctx context.Context, ids []int) (out interface{}, err error) {
	_, err = dao.CmsImage.Ctx(ctx).WhereIn(dao.CmsImage.Columns().Id, ids).Unscoped().Delete()
	if err != nil {
		return nil, err
	}
	return
}

// BackendRecycleBinImageBatchRestore 回收站-图集批量恢复
func (s *sImage) BackendRecycleBinImageBatchRestore(ctx context.Context, ids []int) (out interface{}, err error) {
	_, err = dao.CmsImage.Ctx(ctx).WhereIn(dao.CmsImage.Columns().Id, ids).Unscoped().Update(g.Map{
		dao.CmsImage.Columns().DeletedAt: nil,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s *sImage) buildImagesArr(ctx context.Context, images string) (imagesArr []string, err error) {
	imagesArr = gstr.SplitAndTrim(images, ",")
	return
}

func (s *sImage) buildFlagData(ctx context.Context, flagT, flagR int) (flagData string, err error) {
	var data []string
	if flagT == 1 {
		data = append(data, "t")
	}
	if flagR == 1 {
		data = append(data, "r")
	}
	flagData = gstr.Implode(",", data)
	return
}

func (s *sImage) singleFlag(ctx context.Context, id int, flagType string, targetType string) (out interface{}, err error) {
	m := dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, id)
	var image *entity.CmsImage
	err = m.Scan(&image)
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, gerror.New("数据不存在")
	}
	split := gstr.SplitAndTrim(image.Flag, ",")
	if targetType == "auto" {
		if gstr.InArray(split, flagType) {
			for index, value := range split {
				if value == flagType {
					split = append(split[:index], split[index+1:]...)
					break
				}
			}
		} else {
			split = append(split, flagType)
		}
	} else if targetType == "open" {
		if !gstr.InArray(split, flagType) {
			split = append(split, flagType)
		}
	} else if targetType == "close" {
		if gstr.InArray(split, flagType) {
			for index, value := range split {
				if value == flagType {
					split = append(split[:index], split[index+1:]...)
					break
				}
			}
		}
	} else {
		return nil, gerror.New("操作目的错误")
	}
	flag := gstr.Implode(",", split)
	_, err = m.Data(g.Map{
		dao.CmsImage.Columns().Flag: flag,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sImage) singleStatus(ctx context.Context, id int, targetType string) (out interface{}, err error) {
	m := dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, id)
	var image *entity.CmsImage
	err = m.Scan(&image)
	if err != nil {
		return nil, err
	}
	if image == nil {
		return nil, gerror.New("数据不存在")
	}
	status := 0
	if targetType == "auto" {
		if image.Status == 0 {
			status = 1
		}
	} else if targetType == "open" {
		status = 1
	} else if targetType == "close" {
		status = 0
	} else {
		return nil, gerror.New("操作目的错误")
	}
	_, err = m.Data(g.Map{
		dao.CmsImage.Columns().Status: status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

// BuildThumb 构建图集缩略图
func (s *sImage) BuildThumb(ctx context.Context, in *model.ImageListItem) (out *model.ImageListItem, err error) {
	if len(in.Images) > 0 {
		in.Thumb = service.Util().ImageOrDefaultUrl(in.Images[0])
	} else {
		in.Thumb = service.Util().ImageOrDefaultUrl("")
	}
	out = in
	return
}
