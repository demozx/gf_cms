package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	RecycleBin = cRecycleBin{}
)

type cRecycleBin struct{}

// ArticleList 回收站-文章列表
func (*cRecycleBin) ArticleList(ctx context.Context, req *backendApi.ArticleListReq) (res *backendApi.ArticleListRes, err error) {
	var in *model.ArticleGetListInPut
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	list, err := service.Article().BackendRecycleBinArticleGetList(ctx, in)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "返回成功", list)
	return
}
