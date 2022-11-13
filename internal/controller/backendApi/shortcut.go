package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
)

type cShortcut struct{}

var (
	Shortcut = cShortcut{}
)

// Add 添加快捷方式
func (c *cShortcut) Add(ctx context.Context, req *backendApi.ShortcutAddReq) (res *backendApi.ShortcutAddRes, err error) {
	_, err = service.Shortcut().BackendApiAdd(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "添加成功")
	return
}

// Edit 编辑快捷方式
func (c *cShortcut) Edit(ctx context.Context, req *backendApi.ShortcutEditReq) (res *backendApi.ShortcutEditRes, err error) {
	_, err = service.Shortcut().BackendApiEdit(ctx, req)
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "编辑成功")
	return
}
