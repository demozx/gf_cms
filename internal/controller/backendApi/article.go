package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

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
	sortSlice := make([]*model.ArticleSortMap, 0)
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
