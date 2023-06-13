package pc

import (
	"context"
	"gf_cms/api/pc"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Image = cImage{}
)

type cImage struct{}

// List pc图集列表
func (c *cImage) List(ctx context.Context, req *pc.ImageListReq) (res *pc.ImageListRes, err error) {
	// 栏目详情
	channelInfo, err := Image.channelInfo(ctx, req.ChannelId)
	if err != nil {
		return nil, err
	}
	// 图集列表
	chImagePageList := make(chan *pc.ImageListRes, 1)
	go func() {
		defer close(chImagePageList)
		imagePageList, _ := Image.imagePageList(ctx, req)
		chImagePageList <- imagePageList
	}()
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, gconv.Int(channelInfo.Id))
		chNavigation <- navigation
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		defer close(chTDK)
		tdk, _ := service.Channel().TDK(ctx, channelInfo.Id, 0)
		chTDK <- tdk
	}()
	// 面包屑导航
	chCrumbs := make(chan []*model.ChannelCrumbs, 1)
	go func() {
		defer close(chCrumbs)
		crumbs, _ := service.Channel().Crumbs(ctx, channelInfo.Id)
		chCrumbs <- crumbs
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chGoodsChannelList)
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chTextNewsList)
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		chTextNewsList <- textNewsList
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		defer close(chGuestbookUrl)
		guestbookUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookUrl
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		defer close(chChannelTemplate)
		channelTemplate, _ := service.Channel().PcListTemplate(ctx, channelInfo)
		chChannelTemplate <- channelTemplate
	}()
	imagePageList := <-chImagePageList
	pageInfo := service.PageInfo().PcPageInfo(ctx, imagePageList.Total, imagePageList.Size)
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"channelInfo":         channelInfo,          // 栏目信息
		"navigation":          <-chNavigation,       // 导航
		"tdk":                 <-chTDK,              // TDK
		"crumbs":              <-chCrumbs,           // 面包屑导航
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":        <-chTextNewsList,     // 最新资讯-文字新闻
		"imagePageList":       imagePageList,        // 图集列表
		"pageInfo":            pageInfo,             // 页码
		"guestbookChannelUrl": <-chGuestbookUrl,     // 在线留言栏目url
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// Detail pc图集详情页面
func (c *cImage) Detail(ctx context.Context, req *pc.ImageDetailReq) (res *pc.ImageDetailRes, err error) {
	// 图集详情
	var imageInfo *model.ImageListItem
	err = dao.CmsImage.Ctx(ctx).
		Where(dao.CmsImage.Columns().Id, req.Id).
		Where(dao.CmsImage.Columns().Status, 1).
		Scan(&imageInfo)
	if err != nil {
		return nil, err
	}
	if imageInfo == nil {
		service.Response().Status404(ctx)
	}
	imageInfo.ClickNum++
	// 更新点击量
	go func() {
		//ctx := context.Background()
		_, err = dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, imageInfo.Id).Increment(dao.CmsImage.Columns().ClickNum, 1)
	}()
	// 栏目详情
	channelInfo, err := Image.channelInfo(ctx, imageInfo.ChannelId)
	if err != nil {
		return nil, err
	}
	// 导航栏
	chNavigation := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chNavigation)
		navigation, _ := service.Channel().Navigation(ctx, gconv.Int(channelInfo.Id))
		chNavigation <- navigation
	}()
	// TKD
	chTDK := make(chan *model.ChannelTDK, 1)
	go func() {
		defer close(chTDK)
		tdk, _ := service.Channel().TDK(ctx, gconv.Uint(imageInfo.ChannelId), gconv.Int64(imageInfo.Id))
		chTDK <- tdk
	}()
	// 面包屑导航
	chCrumbs := make(chan []*model.ChannelCrumbs, 1)
	go func() {
		defer close(chCrumbs)
		pcCrumbs, _ := service.Channel().Crumbs(ctx, channelInfo.Id)
		chCrumbs <- pcCrumbs
	}()
	// 产品中心栏目列表
	chGoodsChannelList := make(chan []*model.ChannelNavigationListItem, 1)
	go func() {
		defer close(chGoodsChannelList)
		goodsChannelList, _ := service.Channel().HomeGoodsChannelList(ctx, consts.GoodsChannelId)
		chGoodsChannelList <- goodsChannelList
	}()
	// 最新资讯-文字新闻
	chTextNewsList := make(chan []*model.ArticleListItem, 1)
	go func() {
		defer close(chTextNewsList)
		textNewsList, _ := service.Article().PcHomeTextNewsList(ctx, consts.NewsChannelId)
		chTextNewsList <- textNewsList
	}()
	// 上一篇
	chPrevImage := make(chan *model.ImageLink, 1)
	go func() {
		defer close(chPrevImage)
		prevImage, _ := service.Image().PcPrevImage(ctx, imageInfo.ChannelId, imageInfo.Id)
		chPrevImage <- prevImage
	}()
	// 下一篇
	chNextImage := make(chan *model.ImageLink, 1)
	go func() {
		defer close(chNextImage)
		nextImage, _ := service.Image().PcNextImage(ctx, imageInfo.ChannelId, imageInfo.Id)
		chNextImage <- nextImage
	}()
	// 在线留言栏目链接
	chGuestbookUrl := make(chan string, 1)
	go func() {
		defer close(chGuestbookUrl)
		guestbookChannelUrl, _ := service.GenUrl().ChannelUrl(ctx, consts.GuestbookChannelId, "")
		chGuestbookUrl <- guestbookChannelUrl
	}()
	// 获取模板
	chChannelTemplate := make(chan string, 1)
	go func() {
		defer close(chChannelTemplate)
		channelTemplate, _ := service.Channel().PcDetailTemplate(ctx, channelInfo)
		chChannelTemplate <- channelTemplate
	}()
	err = service.Response().View(ctx, <-chChannelTemplate, g.Map{
		"navigation":          <-chNavigation,       // 导航
		"tdk":                 <-chTDK,              // TDK
		"crumbs":              <-chCrumbs,           // 面包屑导航
		"goodsChannelList":    <-chGoodsChannelList, // 产品中心栏目列表
		"textNewsList":        <-chTextNewsList,     // 最新资讯-文字新闻
		"imageInfo":           imageInfo,            // 图集详情
		"prevImage":           <-chPrevImage,        // 上一篇
		"nextImage":           <-chNextImage,        // 下一篇
		"guestbookChannelUrl": <-chGuestbookUrl,     // 在线留言栏目url
	})
	if err != nil {
		service.Response().Status500(ctx)
		return nil, err
	}
	return
}

// 栏目详情
func (c *cImage) channelInfo(ctx context.Context, channelId int) (out *entity.CmsChannel, err error) {
	err = dao.CmsChannel.Ctx(ctx).
		Where(dao.CmsChannel.Columns().Id, channelId).
		Where(dao.CmsChannel.Columns().Status, 1).
		Where(dao.CmsChannel.Columns().Type, 1).
		Scan(&out)
	if err != nil {
		return
	}
	// 栏目不存在，展示404
	if out == nil {
		service.Response().Status404(ctx)
	}
	return
}

// 获取文章列表分页数据
func (c *cImage) imagePageList(ctx context.Context, in *pc.ImageListReq) (res *pc.ImageListRes, err error) {
	// 当前栏目的所有级别的子栏目id们加自己
	childChannelIds, err := service.Channel().GetChildIds(ctx, in.ChannelId, true)
	if err != nil {
		return nil, err
	}
	m := dao.CmsImage.Ctx(ctx).
		WhereIn(dao.CmsImage.Columns().ChannelId, childChannelIds).
		Where(dao.CmsImage.Columns().Status, 1)
	count, err := m.Count()
	if err != nil {
		return
	}
	var imageList []*model.ImageListItem
	err = m.OrderAsc(dao.CmsImage.Columns().Sort).
		OrderDesc(dao.CmsImage.Columns().Id).
		Page(in.Page, in.Size).
		Scan(&imageList)
	if err != nil {
		return nil, err
	}
	for key, item := range imageList {
		url, err := service.GenUrl().DetailUrl(ctx, consts.ChannelModelImage, gconv.Int(item.Id))
		if err != nil {
			return nil, err
		}
		imageList[key].Router = url
		imageList[key], err = service.Image().BuildThumb(ctx, item)
		if err != nil {
			return nil, err
		}
	}
	res = &pc.ImageListRes{
		List:  imageList,
		Page:  in.Page,
		Size:  in.Size,
		Total: count,
	}
	return
}
