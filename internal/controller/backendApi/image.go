package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 图集模型
type cImage struct{}

var (
	Image = cImage{}
)

func (c *cImage) Index(ctx context.Context, req *backendApi.ImageListReq) (res *backendApi.ImageListRes, err error) {
	var in *model.ImageGetListInPut
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	list, err := service.Image().BackendImageGetList(ctx, in)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "返回成功", list)
	return
}

func (c *cImage) Sort(ctx context.Context, req *backendApi.ImageSortReq) (res *backendApi.ImageSortRes, err error) {
	sortSlice := make([]*model.ImageSortMap, 0, len(req.Sort))
	for _, item := range req.Sort {
		split := gstr.SplitAndTrim(item, "_")
		if len(split) != 2 {
			break
		}
		id := split[0]
		sort := split[1]
		sortData := new(model.ImageSortMap)
		sortData.Id = gvar.New(id).Int()
		sortData.Sort = gvar.New(sort).Int()
		sortSlice = append(sortSlice, sortData)
	}
	_, err = service.Image().Sort(ctx, sortSlice)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cImage) Flag(ctx context.Context, req *backendApi.ImageFlagReq) (res *backendApi.ImageFlagRes, err error) {
	_, err = service.Image().Flag(ctx, req.Ids, req.Flag)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cImage) Status(ctx context.Context, req *backendApi.ImageStatusReq) (res *backendApi.ImageStatusRes, err error) {
	_, err = service.Image().Status(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cImage) Delete(ctx context.Context, req *backendApi.ImageDeleteReq) (res *backendApi.ImageDeleteRes, err error) {
	_, err = service.Image().Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "删除成功", g.Map{})
	return
}

func (c *cImage) Move(ctx context.Context, req *backendApi.ImageMoveReq) (res *backendApi.ImageMoveRes, err error) {
	ids := gstr.Explode(",", req.StrIds)
	_, err = service.Image().Move(ctx, req.ChannelId, ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "移动成功", g.Map{})
	return
}

func (c *cImage) Add(ctx context.Context, req *backendApi.ImageAddReq) (res *backendApi.ImageAddRes, err error) {
	_, err = service.Image().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "添加成功", g.Map{})
	return
}

func (c *cImage) Edit(ctx context.Context, req *backendApi.ImageEditReq) (res *backendApi.ImageEditRes, err error) {
	_, err = service.Image().Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "编辑成功", g.Map{})
	return
}
