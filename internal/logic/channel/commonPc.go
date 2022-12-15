package packed

import (
	"context"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// PcNavigation pc导航
func (s *sChannel) PcNavigation(ctx context.Context, currChannelId int) (out []*model.ChannelPcNavigationListItem, err error) {
	var allOpenChannel []*entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Status, 1).OrderAsc(dao.CmsChannel.Columns().Sort).OrderAsc(dao.CmsChannel.Columns().Id).Scan(&allOpenChannel)
	if err != nil {
		return nil, err
	}
	out, err = Channel().pcNavigationListRecursion(ctx, allOpenChannel, 0, currChannelId)
	return
}

// PcTDK 生成pcTKD
// channelId 栏目id
// detailId  内容页id
func (s *sChannel) PcTDK(ctx context.Context, channelId, detailId int) (out []*model.ChannelPcNavigationListItem, err error) {
	// 列表页TDK
	// 详情页TDK
	return
}

// 递归生成pc导航
func (s *sChannel) pcNavigationListRecursion(ctx context.Context, list []*entity.CmsChannel, pid int, currChannelId int) (out []*model.ChannelPcNavigationListItem, err error) {
	var res []*model.ChannelPcNavigationListItem
	cacheKey := util.PublicCachePreFix + ":pc_navigation_list:pid_" + gconv.String(pid) + "_curr_channel_id_" + gconv.String(currChannelId)
	cached, err := g.Redis().Do(ctx, "GET", cacheKey)
	if err != nil {
		return nil, err
	}
	if !cached.IsEmpty() {
		err := cached.Scan(&res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	// 高亮栏目id
	highlightChannelId := 0
	for _, item := range list {
		var naviItem *model.ChannelPcNavigationListItem
		_ = gconv.Scan(item, &naviItem)
		// 根据频道类型处理url
		switch item.Type {
		case 1:
			// 频道类型
			fallthrough
		case 2:
			// 单页类型
			naviItem.ChannelRouter = item.ListRouter
			if gstr.Contains(item.ListRouter, "{id}") {
				// 如果路由中有{id}，替换id
				naviItem.ChannelRouter, _ = service.GenUrl().PcChannelUrl(ctx, gconv.Int(item.Id), item.ListRouter)
			}
		case 3:
			// 链接类型
			naviItem.ChannelRouter = item.LinkUrl
		default:
			return nil, gerror.New("栏目类型错误")
		}
		// 处理链接打开方式
		naviItem.TriggerType = "_self"
		if item.LinkTrigger == 1 {
			// 新标签打开
			naviItem.TriggerType = "_blank"
		}
		// 判断是否是当前栏目
		if currChannelId > 0 && currChannelId == gconv.Int(naviItem.Id) {
			naviItem.Current = true
			// 顶级栏目高亮
			highlightChannelId = naviItem.Tid
		}
		if item.Pid == pid {
			naviItem.Children, err = Channel().pcNavigationListRecursion(ctx, list, gvar.New(item.Id).Int(), currChannelId)
			if naviItem.Children == nil {
				naviItem.Children = []*model.ChannelPcNavigationListItem{}
			} else {
				naviItem.HasChildren = true
			}
			res = append(res, naviItem)
		}
	}
	if highlightChannelId > 0 {
		// 设置栏目高亮
		for key, item := range res {
			if highlightChannelId == gconv.Int(item.Id) {
				res[key].Highlight = true
			}
		}
	}
	_, err = g.Redis().Do(ctx, "SET", cacheKey, res)
	if err != nil {
		return nil, err
	}
	return res, err
}
