package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
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
