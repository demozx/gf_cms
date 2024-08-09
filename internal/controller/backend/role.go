package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	backendViewAllPermissions, err := service.Permission().BackendViewAll()
	if err != nil {
		return nil, err
	}
	backendApiAllPermissions, err := service.Permission().BackendApiAll()
	if err != nil {
		return nil, err
	}

	listData := list.List
	for key, item := range listData {
		for _key, permission := range item.Permissions {
			v2 := permission.V2 // 权限表中的权限字段
			ruleArr := strings.Split(v2, ".")
			if len(ruleArr) != 3 {
				return nil, gerror.New("权限表中v2字段格式错误")
			}
			for _, _item := range backendViewAllPermissions {
				if ruleArr[0] == _item.Slug {
					for _, _permission := range _item.Permissions {
						if _permission.Slug == ruleArr[1]+"."+ruleArr[2] {
							listData[key].Permissions[_key].Title = _permission.Title
						}
					}
				}
			}
			for _, _item := range backendApiAllPermissions {
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
	backendAllPermissions, err := service.Permission().BackendAll()
	if err != nil {
		return nil, err
	}
	err = service.Response().View(ctx, "backend/role/add.html", g.Map{
		"backendAllPermissions": backendAllPermissions,
	})
	if err != nil {
		return nil, err
	}
	return
}

// Edit 编辑角色
func (c *cRole) Edit(ctx context.Context, req *backend.RoleEditReq) (res *backend.RoleEditRes, err error) {
	backendAllPermissions, err := service.Permission().BackendAll()
	if err != nil {
		return nil, err
	}
	role, err := service.Role().BackendRoleGetOne(ctx, req)
	if err != nil {
		return nil, err
	}
	for key, item := range backendAllPermissions {
		for _key, permission := range item.BackendViewPermissions {
			for _, rolePermission := range role.Permissions {
				if item.Slug+"."+permission.Slug == rolePermission.V2 {
					backendAllPermissions[key].BackendViewPermissions[_key].HasPermission = true
				}
			}
		}
	}
	for key, item := range backendAllPermissions {
		for _key, permission := range item.BackendApiPermissions {
			for _, rolePermission := range role.Permissions {
				if item.Slug+"."+permission.Slug == rolePermission.V2 {
					backendAllPermissions[key].BackendApiPermissions[_key].HasPermission = true
				}
			}
		}
	}
	err = service.Response().View(ctx, "backend/role/edit.html", g.Map{
		"backendAllPermissions": backendAllPermissions,
		"role":                  role,
	})
	if err != nil {
		return nil, err
	}
	return
}
