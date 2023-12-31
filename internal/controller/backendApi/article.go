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

// 文章模型
type cArticle struct{}

var (
	Article = cArticle{}
)

func (c *cArticle) Index(ctx context.Context, req *backendApi.ArticleListReq) (res *backendApi.ArticleListRes, err error) {
	var in *model.ArticleGetListInPut
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	list, err := service.Article().BackendArticleGetList(ctx, in)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "返回成功", list)
	return
}

func (c *cArticle) Sort(ctx context.Context, req *backendApi.ArticleSortReq) (res *backendApi.ArticleSortRes, err error) {
	sortSlice := make([]*model.ArticleSortMap, 0, len(req.Sort))
	for _, item := range req.Sort {
		split := gstr.SplitAndTrim(item, "_")
		if len(split) != 2 {
			break
		}
		id := split[0]
		sort := split[1]
		sortData := new(model.ArticleSortMap)
		sortData.Id = gvar.New(id).Int()
		sortData.Sort = gvar.New(sort).Int()
		sortSlice = append(sortSlice, sortData)
	}
	_, err = service.Article().Sort(ctx, sortSlice)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cArticle) Flag(ctx context.Context, req *backendApi.ArticleFlagReq) (res *backendApi.ArticleFlagRes, err error) {
	_, err = service.Article().Flag(ctx, req.Ids, req.Flag)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cArticle) Status(ctx context.Context, req *backendApi.ArticleStatusReq) (res *backendApi.ArticleStatusRes, err error) {
	_, err = service.Article().Status(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

func (c *cArticle) Delete(ctx context.Context, req *backendApi.ArticleDeleteReq) (res *backendApi.ArticleDeleteRes, err error) {
	_, err = service.Article().Delete(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "删除成功", g.Map{})
	return
}

func (c *cArticle) Move(ctx context.Context, req *backendApi.ArticleMoveReq) (res *backendApi.ArticleMoveRes, err error) {
	ids := gstr.Explode(",", req.StrIds)
	_, err = service.Article().Move(ctx, req.ChannelId, ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "移动成功", g.Map{})
	return
}

func (c *cArticle) Add(ctx context.Context, req *backendApi.ArticleAddReq) (res *backendApi.ArticleAddRes, err error) {
	_, err = service.Article().Add(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "添加成功", g.Map{})
	return
}

func (c *cArticle) Edit(ctx context.Context, req *backendApi.ArticleEditReq) (res *backendApi.ArticleEditRes, err error) {
	_, err = service.Article().Edit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, gcode.CodeOK.Code(), "编辑成功", g.Map{})
	return
}
