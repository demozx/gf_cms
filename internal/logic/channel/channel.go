package channel

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
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
	//for _, channel := range allChannels {
	//	var listItem *model.ChannelBackendApiListItem
	//	if channel.Thumb != "" {
	//		channel.Name = channel.Name + "<span style='color:red;font-size: 12px;margin: 0 2px;'>图</span><span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;id:" + gconv.String(channel.Id) + "</span>"
	//	}
	//	err = gconv.Scan(channel, &listItem)
	//	if err != nil {
	//		return nil, err
	//	}
	//	if channel.Pid == 0 {
	//		channelBackendApiList = append(channelBackendApiList, listItem)
	//	}
	//	if channel.Pid > 0 {
	//		for _key, _item := range channelBackendApiList {
	//			if channel.Pid == _item.Id {
	//				var itemChildren *model.ChannelBackendApiListItem
	//				err = gconv.Scan(listItem, &itemChildren)
	//				if err != nil {
	//					return nil, err
	//				}
	//				channelBackendApiList[_key].Children = append(channelBackendApiList[_key].Children, itemChildren)
	//			}
	//		}
	//	}
	//}
	channelBackendApiList, err = Channel().recursion(channelBackendApiList, allChannels, 0)
	return channelBackendApiList, nil
}

func (*sChannel) recursion(in []*model.ChannelBackendApiListItem, allChannels []*entity.CmsChannel, level int) (out []*model.ChannelBackendApiListItem, err error) {
	//g.Dump("in", in)
	//g.Dump("allChannels", allChannels)
	g.Dump("level", level)
	for _, channel := range allChannels {
		var listItem *model.ChannelBackendApiListItem
		err = gconv.Scan(channel, &listItem)
		if err != nil {
			return nil, err
		}
		if level == 0 && listItem.Level == 0 {
			in = append(in, listItem)
		}
		level += 1
		out, err := Channel().recursion(in, allChannels, level)
		g.Dump("out", out)
		if err != nil {
			return out, err
		}
	}
	g.Dump(in)
	return in, err
}

// ChannelTree 获取栏目分类树
func (*sChannel) ChannelTree(ctx context.Context) (out []*model.ChannelTreeItem, err error) {
	//var allChannels []*entity.CmsChannel
	//err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Status, 1).OrderAsc(dao.CmsChannel.Columns().Sort).Scan(&allChannels)
	//var channelTreeItem []*model.ChannelTreeItem
	//err = gconv.Scan(allChannels, &channelTreeItem)
	//if err != nil {
	//	return nil, err
	//}
	//for key, item := range allChannels {
	//	allChannels[key].Name = "├" + allChannels[key].Name
	//	var emsp = ""
	//	if item.Level > 0 {
	//		for i := 0; i < item.Level; i++ {
	//			emsp += "&emsp;&emsp;"
	//		}
	//	}
	//	allChannels[key].Name = emsp + allChannels[key].Name
	//}
	//err = gconv.Scan(allChannels, &channelTreeItem)
	//if err != nil {
	//	return nil, err
	//}
	//return channelTreeItem, nil

	index, err := service.Channel().Index(ctx)
	if err != nil {
		return nil, err
	}
	g.Dump(index)
	return
}
