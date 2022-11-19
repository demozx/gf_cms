package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/consts"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	Image = cImage{}
)

type cImage struct{}

func (c *cImage) Move(ctx context.Context, req *backend.ImageMoveReq) (res *backend.ImageMoveRes, err error) {
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelImage, 0)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/image/move.html", g.Map{
		"strIds":           req.StrIds,
		"channelModelTree": channelModelTree,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cImage) Add(ctx context.Context, req *backend.ImageAddReq) (res *backend.ImageAddRes, err error) {
	channelId := req.ChannelId
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelImage, channelId)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/image/add.html", g.Map{
		"channelModelTree": channelModelTree,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cImage) Edit(ctx context.Context, req *backend.ImageEditReq) (res *backend.ImageEditRes, err error) {
	var imagesInfo *model.ImageListItem
	err = dao.CmsImage.Ctx(ctx).Where(dao.CmsImage.Columns().Id, req.Id).Scan(&imagesInfo)
	if err != nil {
		return nil, err
	}
	if imagesInfo == nil {
		return nil, gerror.New("图集不存在")
	}
	// 将flag变成数组，方便模板渲染
	flagArr := gstr.Explode(",", imagesInfo.Flag)
	for _, value := range flagArr {
		if value == "r" {
			imagesInfo.FlagR = 1
		}
		if value == "t" {
			imagesInfo.FlagT = 1
		}
	}
	imagesInfo.ImagesStr = gstr.Implode(",", imagesInfo.Images)
	var imagesSrcArr []model.ImagesSrcArrItem
	for _, image := range imagesInfo.Images {
		imagesSrcArr = append(imagesSrcArr, model.ImagesSrcArrItem{
			Src:   image,
			Title: "",
		})
	}
	encodeString, err := gjson.EncodeString(imagesSrcArr)
	if err != nil {
		return nil, err
	}
	imagesInfo.ImagesSrcJson = encodeString
	channelModelTree, err := service.Channel().BackendChannelModelTree(ctx, consts.ChannelModelImage, imagesInfo.ChannelId)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/channel_model/image/edit.html", g.Map{
		"channelModelTree": channelModelTree,
		"imagesInfo":       imagesInfo,
	})
	if err != nil {
		return nil, err
	}
	g.Dump("imagesInfo", imagesInfo)
	return
}
