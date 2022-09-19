package channel

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/util/gconv"
)

type sChannel struct{}

var (
	insChannel = sChannel{}
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
			res = append(res, item)
		}
	}
	return res
}

// BackendChannelTree 获取栏目分类树
func (*sChannel) BackendChannelTree(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error) {
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
	channelBackendApiList = Channel().backendTree(channelBackendApiList)
	return channelBackendApiList, err
}

// 递归生成栏目分类树
func (*sChannel) backendTree(list []*model.ChannelBackendApiListItem) (out []*model.ChannelBackendApiListItem) {
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
		return Channel().backendTree(newList)
	}
	for key, item := range newList {
		var emsp = ""
		for i := 0; i < item.Level; i++ {
			emsp += "&emsp;&emsp;"
		}
		newList[key].Name = emsp + "├&nbsp;" + item.Name
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
	// todo 栏目应该软删除
	// todo 栏目下有内容页不能删除
	_, err = m.Delete()
	if err != nil {
		return nil, err
	}
	return
}
