package channel

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"

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

// Index 获取后台栏目分类接口数据
func (*sChannel) Index(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error) {
	var allChannels []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).OrderAsc(dao.CmsChannel.Columns().Sort).OrderAsc(dao.CmsChannel.Columns().Id).Scan(&allChannels)
	if err != nil {
		return nil, err
	}
	var channelBackendApiList []*model.ChannelBackendApiListItem
	for key, channel := range allChannels {
		if channel.Thumb != "" {
			channel.Name = channel.Name + "<span style='color:red;font-size: 12px;margin: 0 2px;'>图</span><span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;id:" + gconv.String(channel.Id) + "</span>"
			allChannels[key] = channel
		}
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

// ChannelTree 获取栏目分类树
func (*sChannel) ChannelTree(ctx context.Context) (out []*model.ChannelBackendApiListItem, err error) {
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
	channelBackendApiList = Channel().tree(channelBackendApiList)
	return channelBackendApiList, err
}

func (*sChannel) tree(list []*model.ChannelBackendApiListItem) (out []*model.ChannelBackendApiListItem) {
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
		return Channel().tree(newList)
	}
	for key, item := range newList {
		var emsp = ""
		for i := 0; i < item.Level; i++ {
			emsp += "&emsp;&emsp;"
		}
		newList[key].Name = emsp + "├" + item.Name
	}
	return newList
}
