package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

var (
	FriendlyLink = cFriendlyLink{}
)

type cFriendlyLink struct{}

// Status 修改友情链接状态
func (c *cFriendlyLink) Status(ctx context.Context, req *backendApi.FriendlyLinkStatusReq) (res *backendApi.FriendlyLinkStatusRes, err error) {
	_, err = service.FriendlyLink().BackendApiStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefault(ctx)
	return
}

// Add 添加友情链接
func (c *cFriendlyLink) Add(ctx context.Context, req *backendApi.FriendlyLinkAddReq) (res *backendApi.FriendlyLinkAddRes, err error) {
	_, err = service.FriendlyLink().BackendApiAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "添加成功")
	return
}

// Edit 编辑友情链接
func (c *cFriendlyLink) Edit(ctx context.Context, req *backendApi.FriendlyLinkEditReq) (res *backendApi.FriendlyLinkEditRes, err error) {
	_, err = service.FriendlyLink().BackendApiEdit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "编辑成功")
	return
}

// Sort 友情链接排序
func (c *cFriendlyLink) Sort(ctx context.Context, req *backendApi.FriendlyLinkSortReq) (res *backendApi.FriendlyLinkSortRes, err error) {
	_, err = service.FriendlyLink().BackendApiSort(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "排序成功")
	return
}

// BatchDelete 友情链接批量删除
func (c *cFriendlyLink) BatchDelete(ctx context.Context, req *backendApi.FriendlyLinkBatchDeleteReq) (res *backendApi.FriendlyLinkBatchDeleteRes, err error) {
	_, err = service.FriendlyLink().BackendApiBatchDelete(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "删除成功")
	return
}
