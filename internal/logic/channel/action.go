package packed

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/errors/gerror"
)

// HomeAboutChannel 关于我们
func (s *sChannel) HomeAboutChannel(ctx context.Context, channelId int) (channel *entity.CmsChannel, err error) {
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channel)
	if err != nil {
		return nil, err
	}
	if channel == nil {
		return nil, gerror.New("栏目不存在")
	}
	return
}

// HomeGoodsChannelList 首页产品栏目列表
func (s *sChannel) HomeGoodsChannelList(ctx context.Context, channelId int) (out []*model.ChannelNavigationListItem, err error) {
	var list []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).
		Where(dao.CmsChannel.Columns().Tid, channelId).
		Where(dao.CmsChannel.Columns().Status, 1).
		OrderAsc(dao.CmsChannel.Columns().Sort).
		OrderDesc(dao.CmsChannel.Columns().Id).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, gerror.New("栏目数据不存在")
	}
	res, err := Channel().navigationListRecursion(ctx, list, channelId, 0)
	if err != nil {
		return nil, err
	}
	out = res
	return
}
