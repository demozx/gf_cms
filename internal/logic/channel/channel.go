package channel

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/util/gconv"
)

type sChannel struct{}

var (
	insChannel = sChannel{}
)

const (
	ModelArticle     = "article"
	ModelArticleDesc = "文章"
)

func init() {
	service.RegisterChannel(New())
}

func New() *sChannel {
	return &sChannel{}
}

func Channel() *sChannel {
	return &insChannel
}

// BackendIndex 获取后台栏目分类接口数据
func (*sChannel) BackendIndex(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error) {
	var allChannels []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).OrderAsc(dao.CmsChannel.Columns().Sort).OrderAsc(dao.CmsChannel.Columns().Id).Scan(&allChannels)
	if err != nil {
		return nil, err
	}
	var channelBackendApiList []*model.ChannelBackendApiListItem
	for key, channel := range allChannels {
		if channel.Thumb != "" {
			channel.Name = channel.Name + "<span style='color:red;font-size: 12px;margin: 0 2px;'>图</span><span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;id:" + gconv.String(channel.Id) + "</span>"

		} else {
			channel.Name = channel.Name + "<span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;id:" + gconv.String(channel.Id) + "</span>"
		}
		allChannels[key] = channel
	}
	err = gconv.Scan(allChannels, &channelBackendApiList)
	if err != nil {
		return nil, err
	}
	channelBackendApiList = Channel().recursion(channelBackendApiList, 0)
	return channelBackendApiList, nil
}

func (*sChannel) recursion(list []*model.ChannelBackendApiListItem, pid int) (out []*model.ChannelBackendApiListItem) {
	res := make([]*model.ChannelBackendApiListItem, 0)
	for _, item := range list {
		if item.Pid == pid {
			item.Children = Channel().recursion(list, item.Id)
			if item.Children == nil {
				item.Children = make([]*model.ChannelBackendApiListItem, 0)
			}
			item.ModelDesc = service.Channel().BackendModelDesc(item.Model)
			res = append(res, item)
		}
	}
	return res
}

// BackendChannelTree 获取栏目分类树
func (*sChannel) BackendChannelTree(ctx context.Context, channelId int) (out []*model.ChannelBackendApiListItem, err error) {
	var allChannels []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).OrderAsc(dao.CmsChannel.Columns().Sort).OrderAsc(dao.CmsChannel.Columns().Id).Scan(&allChannels)
	if err != nil {
		return nil, err
	}
	var channelBackendApiList []*model.ChannelBackendApiListItem
	err = gconv.Scan(allChannels, &channelBackendApiList)
	if err != nil {
		return nil, err
	}
	channelBackendApiList = Channel().recursion(channelBackendApiList, 0)
	//g.Dump(channelBackendApiList)
	channelBackendApiList = Channel().backendTree(channelBackendApiList, channelId)

	return channelBackendApiList, err
}

// 递归生成栏目分类树
func (*sChannel) backendTree(list []*model.ChannelBackendApiListItem, selectedPid int) (out []*model.ChannelBackendApiListItem) {
	var hasChildren = false
	newList := make([]*model.ChannelBackendApiListItem, 0)
	for _, item := range list {
		newItem := new(model.ChannelBackendApiListItem)
		newItem.Id = item.Id
		newItem.Pid = item.Pid
		newItem.Level = item.Level
		newItem.Status = item.Status
		newItem.Name = item.Name
		newItem.Children = nil
		newList = append(newList, newItem)
		if len(item.Children) > 0 {
			hasChildren = true
			for _, childrenItem := range item.Children {
				newList = append(newList, childrenItem)
			}
		}
	}
	if hasChildren == true {
		return Channel().backendTree(newList, selectedPid)
	}
	for key, item := range newList {
		var emsp = ""
		for i := 0; i < item.Level; i++ {
			emsp += "&emsp;&emsp;"
		}
		newList[key].Name = emsp + "├" + item.Name
		if item.Id == selectedPid {
			newList[key].Selected = 1
		}
	}
	return newList
}

// BackendApiStatus 状态禁用启用
func (*sChannel) BackendApiStatus(ctx context.Context, in *backendApi.ChannelStatusApiReq) (out *backendApi.ChannelStatusApiRes, err error) {
	var first *entity.CmsChannel
	var m = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Id)
	err = m.Scan(&first)
	if err != nil {
		return nil, err
	}
	if first == nil {
		return nil, gerror.New("栏目不存在")
	}
	status := 1
	if first.Status == 1 {
		status = 0
	}
	_, err = m.Data(g.Map{
		dao.CmsChannel.Columns().Status: status,
	}).Update()
	if err != nil {
		return nil, err
	}
	return
}

func (*sChannel) BackendApiDelete(ctx context.Context, in *backendApi.ChannelDeleteApiReq) (out *backendApi.ChannelDeleteApiRes, err error) {
	var first *entity.CmsChannel
	var m = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Id)
	err = m.Scan(&first)
	if err != nil {
		return nil, err
	}
	if first == nil {
		return nil, gerror.New("栏目不存在")
	}
	var children *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Pid, in.Id).Scan(&children)
	if err != nil {
		return nil, err
	}
	if children != nil {
		return nil, gerror.New("当前栏目下有子栏目，不允许删除")
	}
	// todo 栏目下有内容页不能删除
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	return
}

func (*sChannel) BackendApiAdd(ctx context.Context, in *backendApi.ChannelAddApiReq) (out *backendApi.ChannelAddApiRes, err error) {
	var entityData *entity.CmsChannel
	err = gconv.Scan(in, &entityData)
	if err != nil {
		return nil, err
	}
	var parent *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Pid).Scan(&parent)
	if err != nil {
		return nil, err
	}
	entityData.Level = 0
	if parent != nil {
		entityData.Level = parent.Level + 1
	}
	_, err = dao.CmsChannel.Ctx(ctx).Data(entityData).Insert()
	if err != nil {
		return nil, err
	}
	return
}

func (*sChannel) BackendApiEdit(ctx context.Context, in *backendApi.ChannelEditApiReq) (out *backendApi.ChannelEditApiRes, err error) {
	var oldLevel = 0
	var newLevel = 0
	var newPid = 0
	if in.Pid > 0 {
		parent, err := Channel().GetOneById(ctx, in.Pid)
		if err != nil {
			return nil, err
		}
		oldLevel = parent.Level
		newLevel = parent.Level + 1
		newPid = gvar.New(parent.Id).Int()
	}
	channel, err := Channel().GetOneById(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if oldLevel >= channel.Level {
		return nil, gerror.New("不能选择同级别或级别在自己下边的分类")
	}
	if in.Pid == in.Id {
		return nil, gerror.New("自己不能是自己的父级分类")
	}
	var data *entity.CmsChannel
	err = gconv.Scan(in, &data)
	if err != nil {
		return nil, err
	}
	data.Level = newLevel
	data.Pid = newPid
	_, err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, in.Id).Update(data)
	if err != nil {
		return nil, err
	}
	return
}

func (*sChannel) GetOneById(ctx context.Context, id int) (out *entity.CmsChannel, err error) {
	var channel *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, id).Scan(&channel)
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, gerror.New("栏目分类不存在")
	}
	return channel, nil
}

func (*sChannel) BackendModelMap() map[string]string {
	var modelMap = make(map[string]string)
	modelMap[ModelArticle] = ModelArticleDesc
	return modelMap
}

func (*sChannel) BackendModelDesc(model string) string {
	modelMap := Channel().BackendModelMap()
	out, isOk := modelMap[model]
	if isOk == false {
		return ""
	}
	return out
}
