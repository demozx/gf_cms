package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
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

// ArticleBatchDestroy 回收站-文章批量永久删除
func (*cRecycleBin) ArticleBatchDestroy(ctx context.Context, req *backendApi.ArticleBatchDestroyReq) (res *backendApi.ArticleBatchDestroyRes, err error) {
	_, err = service.Article().BackendRecycleBinArticleBatchDestroy(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "删除成功", g.Map{})
	return
}

// ArticleBatchRestore 回收站-文章批量恢复
func (*cRecycleBin) ArticleBatchRestore(ctx context.Context, req *backendApi.ArticleBatchRestoreReq) (res *backendApi.ArticleBatchRestoreRes, err error) {
	_, err = service.Article().BackendRecycleBinArticleBatchRestore(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "恢复成功", g.Map{})
	return
}

// ImageList 回收站-图集列表
func (*cRecycleBin) ImageList(ctx context.Context, req *backendApi.ImageListReq) (res *backendApi.ImageListRes, err error) {
	var in *model.ImageGetListInPut
	err = gconv.Scan(req, &in)
	if err != nil {
		return nil, err
	}
	list, err := service.Image().BackendRecycleBinImageGetList(ctx, in)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "返回成功", list)
	return
}

// ImageBatchDestroy 回收站-文章批量永久删除
func (*cRecycleBin) ImageBatchDestroy(ctx context.Context, req *backendApi.ImageBatchDestroyReq) (res *backendApi.ImageBatchDestroyRes, err error) {
	_, err = service.Image().BackendRecycleBinImageBatchDestroy(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "删除成功", g.Map{})
	return
}

// ImageBatchRestore 回收站-文章批量恢复
func (*cRecycleBin) ImageBatchRestore(ctx context.Context, req *backendApi.ImageBatchRestoreReq) (res *backendApi.ImageBatchRestoreRes, err error) {
	_, err = service.Image().BackendRecycleBinImageBatchRestore(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJson(ctx, service.Response().SuccessCodeDefault(), "恢复成功", g.Map{})
	return
}
