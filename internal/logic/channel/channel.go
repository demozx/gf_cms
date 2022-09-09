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
	for _, channel := range allChannels {
		var listItem *model.ChannelBackendApiListItem
		if channel.Thumb != "" {
			channel.Name = channel.Name + "<span style='color:red;font-size: 12px;margin: 0 2px;'>图</span><span id='id_" + gconv.String(channel.Id) + "' class='cate_id'>&nbsp;id:" + gconv.String(channel.Id) + "</span>"
		}
		err = gconv.Scan(channel, &listItem)
		if err != nil {
			return nil, err
		}
		if channel.Pid == 0 {
			channelBackendApiList = append(channelBackendApiList, listItem)
		}
		if channel.Pid > 0 {
			for _key, _item := range channelBackendApiList {
				if channel.Pid == _item.Id {
					var itemChildren *model.ChannelBackendApiListItemChildren
					err = gconv.Scan(listItem, &itemChildren)
					if err != nil {
						return nil, err
					}
					channelBackendApiList[_key].Children = append(channelBackendApiList[_key].Children, itemChildren)
				}
			}
		}
	}
	return channelBackendApiList, nil
}
