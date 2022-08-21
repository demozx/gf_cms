package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

var (
	Role = cRole{}
)

type cRole struct{}

// Index 角色列表
func (c *cRole) Index(ctx context.Context, req *backend.RoleIndexReq) (res *backend.RoleIndexRes, err error) {
	list, err := service.Role().BackendRoleGetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	backendAllPermissions := service.Permission().BackendAll()

	listData := list.List
	for key, item := range listData {
		for _key, permission := range item.Permissions {
			v2 := permission.V2 // 权限表中的权限字段
			ruleArr := strings.Split(v2, ".")
			if len(ruleArr) != 3 {
				return nil, gerror.New("权限表中v2字段格式错误")
			}
			for _, _item := range backendAllPermissions {
				if ruleArr[0] == _item.Slug {
					for _, _permission := range _item.Permissions {
						if _permission.Slug == ruleArr[1]+"."+ruleArr[2] {
							listData[key].Permissions[_key].Title = _permission.Title
						}
					}
				}
			}
		}
	}
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/role/index.html", g.Map{
		"list":     list,
		"pageInfo": service.PageInfo().LayUiPageInfo(ctx, list.Total, list.Size),
	})
	if err != nil {
		return nil, err
	}
	return
}

// Add 添加角色
func (c *cRole) Add(ctx context.Context, req *backend.RoleAddReq) (res *backend.RoleAddRes, err error) {
	backendAllPermission := service.Permission().BackendAll()
	err = service.Response().View(ctx, "backend/role/add.html", g.Map{
		"backendAllPermission": backendAllPermission,
	})
	if err != nil {
		return nil, err
	}
	return
}

// Edit 编辑角色
func (c *cRole) Edit(ctx context.Context, req *backend.RoleEditReq) (res *backend.RoleEditRes, err error) {
	backendAllPermission := service.Permission().BackendAll()
	role, err := service.Role().BackendRoleGetOne(ctx, req)
	//g.Dump(role, backendAllPermission)
	if err != nil {
		return nil, err
	}
	for key, item := range backendAllPermission {
		for _key, permission := range item.Permissions {
			for _, rolePermission := range role.Permissions {
				if item.Slug+"."+permission.Slug == rolePermission.V2 {
					backendAllPermission[key].Permissions[_key].HasPermission = true
				}
			}
		}
	}
	err = service.Response().View(ctx, "backend/role/edit.html", g.Map{
		"backendAllPermission": backendAllPermission,
		"role":                 role,
	})
	if err != nil {
		return nil, err
	}
	return
}
