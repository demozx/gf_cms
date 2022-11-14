package shortcut

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/api/backendApi"
	"gf_cms/internal/dao"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

var (
	insShortcut = sShortcut{}
)

type sShortcut struct{}

func init() {
	service.RegisterShortcut(New())
}

func New() *sShortcut {
	return &sShortcut{}
}

func Shortcut() *sShortcut {
	return &insShortcut
}

// BackendIndex 列表
func (s *sShortcut) BackendIndex(ctx context.Context) (out []*entity.CmsShortcut, err error) {
	cmsAdmin, err := Shortcut().backendGetUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	var shortcutList []*entity.CmsShortcut
	err = dao.CmsShortcut.Ctx(ctx).Where(dao.CmsShortcut.Columns().AccountId, cmsAdmin.Id).OrderAsc(dao.CmsShortcut.Columns().Sort).OrderAsc(dao.CmsShortcut.Columns().Id).Scan(&shortcutList)
	if err != nil {
		return nil, err
	}
	for key, item := range shortcutList {
		shortcutList[key].Route = service.ViewBindFun().BackendUrl(item.Route)
	}
	out = shortcutList
	return
}

// BackendEdit 后台编辑返回信息
func (s *sShortcut) BackendEdit(ctx context.Context, in *backend.ShortcutEditReq) (out *entity.CmsShortcut, err error) {
	cmsAdmin, err := Shortcut().backendGetUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	var cmsShortcut *entity.CmsShortcut
	err = dao.CmsShortcut.Ctx(ctx).Where(dao.CmsShortcut.Columns().Id, in.Id).Where(dao.CmsShortcut.Columns().AccountId, cmsAdmin.Id).Scan(&cmsShortcut)
	if err != nil {
		return nil, err
	}
	if cmsShortcut == nil {
		return nil, gerror.New("快捷方式不存在")
	}
	out = cmsShortcut
	return
}

// BackendApiAdd 添加快捷方式
func (s *sShortcut) BackendApiAdd(ctx context.Context, in *backendApi.ShortcutAddReq) (out interface{}, err error) {
	split := gstr.SplitAndTrim(in.Route, "/")
	countI := gstr.CountI(in.Route, "/")
	if len(split) != 2 || countI != 2 {
		return nil, gerror.New("路由格式错误")
	}
	cmsAdmin, err := Shortcut().backendGetUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	one, err := dao.CmsShortcut.Ctx(ctx).Where(dao.CmsShortcut.Columns().AccountId, cmsAdmin.Id).Where(dao.CmsShortcut.Columns().Route, in.Route).One()
	if err != nil {
		return nil, err
	}
	if !one.IsEmpty() {
		return nil, gerror.New("路由已存在")
	}
	_, err = dao.CmsShortcut.Ctx(ctx).Insert(g.Map{
		"account_id": cmsAdmin.Id,
		"name":       in.Name,
		"route":      in.Route,
	})
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiEdit 编辑快捷方式
func (s *sShortcut) BackendApiEdit(ctx context.Context, in *backendApi.ShortcutEditReq) (out interface{}, err error) {
	split := gstr.SplitAndTrim(in.Route, "/")
	countI := gstr.CountI(in.Route, "/")
	if len(split) != 2 || countI != 2 {
		return nil, gerror.New("路由格式错误")
	}
	cmsAdmin, err := Shortcut().backendGetUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	one, err := dao.CmsShortcut.Ctx(ctx).
		Where(dao.CmsShortcut.Columns().AccountId, cmsAdmin.Id).
		Where(dao.CmsShortcut.Columns().Route, in.Route).
		WhereNot(dao.CmsShortcut.Columns().Id, in.Id).
		One()
	if err != nil {
		return nil, err
	}
	if !one.IsEmpty() {
		return nil, gerror.New("路由已存在")
	}
	affected, err := dao.CmsShortcut.Ctx(ctx).
		Where(dao.CmsShortcut.Columns().Id, in.Id).
		Where(dao.CmsShortcut.Columns().AccountId, cmsAdmin.Id).
		Data(in).
		UpdateAndGetAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, gerror.New("路由不存在")
	}
	return
}

// BackendApiBatchDelete 删除快捷方式
func (s *sShortcut) BackendApiBatchDelete(ctx context.Context, in *backendApi.ShortcutBatchDeleteReq) (out interface{}, err error) {
	cmsAdmin, err := Shortcut().backendGetUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	_, err = dao.CmsShortcut.Ctx(ctx).
		WhereIn(dao.CmsShortcut.Columns().Id, in.Ids).
		Where(dao.CmsShortcut.Columns().AccountId, cmsAdmin.Id).
		Delete()
	if err != nil {
		return nil, err
	}
	return
}

// BackendApiSort 排序
func (s *sShortcut) BackendApiSort(ctx context.Context, in *backendApi.ShortcutSortReq) (out interface{}, err error) {
	var data []*model.ShortcutSortItem
	for _, item := range in.Sort {
		split := gstr.SplitAndTrim(item, "_")
		if len(split) != 2 {
			continue
		}
		shortcutSortItem := &model.ShortcutSortItem{
			Id:   gvar.New(split[0]).Int(),
			Sort: gvar.New(split[1]).Int(),
		}
		data = append(data, shortcutSortItem)
	}
	_, err = dao.CmsShortcut.Ctx(ctx).Data(data).Save()
	if err != nil {
		return nil, err
	}
	return
}

// 从session获取当前登录用户
func (s *sShortcut) backendGetUserFromSession(ctx context.Context) (out *entity.CmsAdmin, err error) {
	cmsAdmin, err := service.Permission().BackendGetUserFromSession(ctx)
	if err != nil {
		return nil, err
	}
	return cmsAdmin, nil
}
