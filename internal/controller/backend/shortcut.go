package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Shortcut = cShortcut{}
)

type cShortcut struct{}

// Index 快捷方式列表
func (c *cShortcut) Index(ctx context.Context, req *backend.ShortcutIndexReq) (res *backend.ShortcutIndexRes, err error) {
	list, err := service.Shortcut().BackendIndex(ctx)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/shortcut/index.html", g.Map{
		"list": list,
	})
	if err != nil {
		return nil, err
	}
	return
}

// Edit 编辑快捷方式
func (c *cShortcut) Edit(ctx context.Context, req *backend.ShortcutEditReq) (res *backend.ShortcutEditRes, err error) {
	shortcut, err := service.Shortcut().BackendEdit(ctx, req)
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "/backend/shortcut/edit.html", g.Map{
		"shortcut": shortcut,
	})
	if err != nil {
		return nil, err
	}
	return
}
