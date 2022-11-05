package adChannel

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

type (
	sAdChannel struct{}
)

var (
	insAdChannel = sAdChannel{}
)

func New() *sAdChannel {
	return &sAdChannel{}
}

func Admin() *sAdChannel {
	return &insAdChannel
}

func init() {
	service.RegisterAdChannel(New())
}

func (s *sAdChannel) Add(ctx context.Context, in *backendApi.AdChannelAddReq) (out interface{}, err error) {
	id, err := dao.CmsAdChannel.Ctx(ctx).Where(dao.CmsAdChannel.Columns().ChannelName, in.ChannelName).Value(dao.CmsAdChannel.Columns().Id)
	if err != nil {
		return nil, err
	}
	if !id.IsEmpty() {
		return nil, gerror.New("分类名称已存在")
	}
	_, err = dao.CmsAdChannel.Ctx(ctx).Insert(g.Map{
		dao.CmsAdChannel.Columns().ChannelName: in.ChannelName,
		dao.CmsAdChannel.Columns().Remarks:     in.Remarks,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s *sAdChannel) Edit(ctx context.Context, in *backendApi.AdChannelEditReq) (out interface{}, err error) {
	id, err := dao.CmsAdChannel.Ctx(ctx).WhereNot(dao.CmsAdChannel.Columns().Id, in.ChannelId).
		Where(dao.CmsAdChannel.Columns().ChannelName, in.ChannelName).
		Value(dao.CmsAdChannel.Columns().Id)
	if err != nil {
		return nil, err
	}
	if !id.IsEmpty() {
		return nil, gerror.New("分类名称已存在")
	}
	id, err = dao.CmsAdChannel.Ctx(ctx).Where(dao.CmsAdChannel.Columns().Id, in.ChannelId).Value(dao.CmsAdChannel.Columns().Id)
	if err != nil {
		return nil, err
	}
	if id.IsEmpty() {
		return nil, gerror.New("分类信息不存在")
	}
	_, err = dao.CmsAdChannel.Ctx(ctx).Where(dao.CmsAdChannel.Columns().Id, in.ChannelId).Update(g.Map{
		dao.CmsAdChannel.Columns().ChannelName: in.ChannelName,
		dao.CmsAdChannel.Columns().Remarks:     in.Remarks,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s *sAdChannel) Delete(ctx context.Context, in *backendApi.AdChannelDeleteReq) (out interface{}, err error) {
	_, err = dao.CmsAdChannel.Ctx(ctx).Where(dao.CmsAdChannel.Columns().Id, in.ChannelId).Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (s *sAdChannel) Sort(ctx context.Context, in *backendApi.AdChannelSortReq) (out interface{}, err error) {
	sortSlice := make([]*model.AdChannelSortMap, 0)
	for _, value := range in.Sort {
		split := gstr.SplitAndTrim(value, "_")
		if len(split) != 2 {
			break
		}
		id := split[0]
		sort := split[1]
		sortData := new(model.AdChannelSortMap)
		sortData.Id = gvar.New(id).Int()
		sortData.Sort = gvar.New(sort).Int()
		sortSlice = append(sortSlice, sortData)
		_, err = dao.CmsAdChannel.Ctx(ctx).Save(sortSlice)
		if err != nil {
			return nil, err
		}
	}
	return
}
