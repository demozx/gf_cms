package packed

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/ghtml"
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

// PcTDK 生成pcTDK
// channelId 栏目id
// detailId  内容页id
func (s *sChannel) PcTDK(ctx context.Context, channelId uint, detailId int64) (out *model.ChannelTDK, err error) {
	// 首页
	if channelId == 0 {
		out = &model.ChannelTDK{
			Title:       service.Util().GetSetting("web_name"),
			Description: service.Util().GetSetting("description"),
			Keywords:    service.Util().GetSetting("keywords"),
		}
		return
	}
	cacheKey := util.PublicCachePreFix + ":pc_tdk:channel_" + gconv.String(channelId) + "_detail_" + gconv.String(detailId)
	cached, err := g.Redis().Do(ctx, "GET", cacheKey)
	if err != nil {
		return nil, err
	}
	if !cached.IsEmpty() {
		err := cached.Scan(&out)
		if err != nil {
			return nil, err
		}
		return out, nil
	}
	var channelInfo *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channelInfo)
	if err != nil {
		return nil, err
	}
	if channelInfo == nil {
		return nil, gerror.New("栏目不存在")
	}
	title := ""
	description := ""
	keywords := ""
	if detailId == 0 {
		description = channelInfo.Description
	}
	title, err = Channel().pcChannelTitleRecursion(ctx, gconv.Uint(channelInfo.Pid), channelInfo.Name)
	if err != nil {
		return nil, err
	}
	// 有详情页，比如文章详情、图集详情，title要拼接上详情页的title，description要使用详情页的，keyword要使用详情页的
	if detailId > 0 {
		detailInfo, err := service.ChannelModel().GetDetailOneByChannelId(ctx, channelInfo.Id, detailId)
		if err != nil {
			return nil, err
		}
		switch channelInfo.Model {
		case consts.ChannelModelArticle:
			var article *entity.CmsArticle
			err := gconv.Scan(detailInfo, &article)
			if err != nil {
				return nil, err
			}
			title = article.Title + "_" + title
			keywords = article.Keyword
			description = article.Description
		case consts.ChannelModelImage:
			var image *entity.CmsImage
			err := gconv.Scan(detailInfo, &image)
			if err != nil {
				return nil, err
			}
			title = image.Title + "_" + title
			description = image.Description
		}
	}
	description = gstr.SubStrRune(ghtml.StripTags(description), 0, 255)
	out = &model.ChannelTDK{
		Title:       title,
		Keywords:    keywords,
		Description: description,
	}
	_, err = g.Redis().Do(ctx, "SET", cacheKey, out)
	if err != nil {
		return nil, err
	}
	return
}

// PcCrumbs 生成pc面包屑导航
// channelId 栏目id
// detailId  内容页id
func (s *sChannel) PcCrumbs(ctx context.Context, channelId uint) (out []*model.ChannelCrumbs, err error) {
	cacheKey := util.PublicCachePreFix + ":pc_crumbs:channel_" + gconv.String(channelId)
	cached, err := g.Redis().Do(ctx, "GET", cacheKey)
	if err != nil {
		return nil, err
	}
	if !cached.IsEmpty() {
		err := cached.Scan(&out)
		if err != nil {
			return nil, err
		}
		return out, nil
	}
	out, err = Channel().pcCrumbsRecursion(ctx, channelId, nil)
	if err != nil {
		return nil, err
	}
	_, err = g.Redis().Do(ctx, "SET", cacheKey, out)
	return
}

// PcListTemplate 获取pc栏目列表模板
func (s *sChannel) PcListTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error) {
	switch channel.Type {
	case 1:
		// 频道
		switch channel.Model {
		case consts.ChannelModelArticle:
			// 文章
			template = "/pc/article/list.html"
		case consts.ChannelModelImage:
			// 图集
			template = "/pc/image/list.html"
		}
	case 2:
		// 单页
		template = "/pc/single_page/detail.html"
	}
	if len(channel.ListTemplate) > 0 {
		// 后台配置的时候不需要加“/pc”，程序自动找指定模块下的模板
		template = channel.ListTemplate
		if gstr.HasPrefix(template, "/pc") {
			template = "/pc" + template
		}
	}
	return
}

// PcDetailTemplate 获取pc栏目详情模板
func (s *sChannel) PcDetailTemplate(ctx context.Context, channel *entity.CmsChannel) (template string, err error) {
	switch channel.Model {
	case consts.ChannelModelArticle:
		// 文章
		template = "/pc/article/detail.html"
	case consts.ChannelModelImage:
		// 图集
		template = "/pc/image/detail.html"
	}
	if len(channel.DetailTemplate) > 0 {
		// 后台配置的时候不需要加“/pc”，程序自动找指定模块下的模板
		template = channel.DetailTemplate
		if gstr.HasPrefix(template, "/pc") {
			template = "/pc" + template
		}
	}
	return
}

// 递归生成面包屑导航
func (s *sChannel) pcCrumbsRecursion(ctx context.Context, channelId uint, crumbs []*model.ChannelCrumbs) (out []*model.ChannelCrumbs, err error) {
	if channelId == 0 {
		return crumbs, nil
	}
	var channelInfo *model.ChannelPcNavigationListItem
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelId).Scan(&channelInfo)
	if err != nil {
		return nil, err
	}
	if channelInfo == nil {
		return nil, gerror.New("栏目不存在")
	}
	channelRouter, err := service.GenUrl().ChannelUrl(ctx, gconv.Int(channelInfo.Id), "")
	if err != nil {
		return nil, err
	}
	crumbs = append(crumbs, &model.ChannelCrumbs{
		Name:   channelInfo.Name,
		Router: channelRouter,
	})
	// 将原面包屑切片倒叙
	var invertedCrumbs = make([]*model.ChannelCrumbs, 0, len(crumbs))
	for key, _ := range crumbs {
		index := len(crumbs) - key - 1
		invertedCrumbs = append(invertedCrumbs, crumbs[index])
	}
	return Channel().pcCrumbsRecursion(ctx, gconv.Uint(channelInfo.Pid), invertedCrumbs)
}

// pc栏目title递归组成（仅栏目，不含内容详情页的title）
func (s *sChannel) pcChannelTitleRecursion(ctx context.Context, channelPid uint, title string) (out string, err error) {
	// 顶级，返回
	if channelPid == 0 {
		return title + "-" + service.Util().GetSetting("web_name"), nil
	}
	var channelInfo *entity.CmsChannel
	err = dao.CmsChannel.Ctx(ctx).Where(dao.CmsChannel.Columns().Id, channelPid).Scan(&channelInfo)
	if err != nil {
		return "", err
	}
	if channelInfo == nil {
		return "", gerror.New("栏目不存在")
	}
	title = title + "_" + channelInfo.Name
	out, err = Channel().pcChannelTitleRecursion(ctx, gconv.Uint(channelInfo.Pid), title)
	if err != nil {
		return "", err
	}
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
				naviItem.ChannelRouter, _ = service.GenUrl().ChannelUrl(ctx, gconv.Int(item.Id), item.ListRouter)
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
