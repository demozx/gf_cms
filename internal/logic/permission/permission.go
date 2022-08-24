package permission

import (
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/admin"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"log"
	"os"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
)

type sPermission struct{}

var (
	insPermission = sPermission{}
)

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// Permission 权限
func Permission() *sPermission {
	return &insPermission
}

func (*sPermission) readYamlConfig(path string) (*model.PermissionConfig, error) {
	conf := &model.PermissionConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	//fmt.Println("conf: ", conf)
	return conf, nil
}

func (*sPermission) readYaml() *model.PermissionConfig {
	conf, err := Permission().readYamlConfig(util.Util().SystemRoot() + "/manifest/config/permission.yaml")
	//fmt.Println(conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendAll 获取后台全部权限（view和api）
func (*sPermission) BackendAll() []model.PermissionAllItem {
	backendViewAll := service.Permission().BackendViewAll()
	backendApiAll := service.Permission().BackendApiAll()
	var permissionAll []model.PermissionAllItem
	for _, viewItem := range backendViewAll {
		for _, apiItem := range backendApiAll {
			permissionAllItem := model.PermissionAllItem{}
			permissionAllItem.Title = viewItem.Title
			permissionAllItem.Slug = viewItem.Slug
			for _, viewItemPermission := range viewItem.Permissions {
				permissionAllItem.BackendViewPermissions = append(permissionAllItem.BackendViewPermissions, viewItemPermission)
			}
			if viewItem.Slug == apiItem.Slug {
				for _, apiItemPermission := range apiItem.Permissions {
					permissionAllItem.BackendApiPermissions = append(permissionAllItem.BackendApiPermissions, apiItemPermission)
				}
			}
			permissionAll = append(permissionAll, permissionAllItem)
		}
	}
	return permissionAll
}

// BackendViewAll Backend 获取全部后台权限
func (*sPermission) BackendViewAll() []model.PermissionGroups {
	cacheKey := util.PublicCachePreFix + ":permissions:backend_view_all"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var permissionGroups []model.PermissionGroups
		if err = result.Structs(&permissionGroups); err != nil {
			panic(err)
		}
		return permissionGroups
	}
	backendViewAll := Permission().readYaml().BackendView.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendViewAll)
	if err != nil {
		panic(err)
	}
	return backendViewAll
}

// BackendApiAll Backend 获取全部后台接口权限
func (*sPermission) BackendApiAll() []model.PermissionGroups {
	cacheKey := util.PublicCachePreFix + ":permissions:backend_api_all"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var permissionGroups []model.PermissionGroups
		if err = result.Structs(&permissionGroups); err != nil {
			panic(err)
		}
		return permissionGroups
	}
	backendApiAll := Permission().readYaml().BackendApi.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendApiAll)
	if err != nil {
		panic(err)
	}
	return backendApiAll
}

// BackendMyView 获取我的所有后台视图权限
func (*sPermission) BackendMyView(accountId string) []gdb.Value {
	roleIds := admin.Admin().GetRoleIdsByAccountId(accountId)
	if len(roleIds) == 0 {
		panic("用户无任何角色")
	}
	myPermissions, err := dao.CmsRulePermissions.Ctx(util.Ctx).
		Where("p_type", "p").
		WhereIn("v0", roleIds).
		Where("v1", "backend").
		Fields("v2").
		Array()
	if err != nil {
		panic(err)
	}
	return myPermissions
}

// BackendMyApi 获取我的所有后台接口权限
func (*sPermission) BackendMyApi(accountId string) []gdb.Value {
	roleIds := admin.Admin().GetRoleIdsByAccountId(accountId)
	if len(roleIds) == 0 {
		panic("用户无任何角色")
	}
	myPermissions, err := dao.CmsRulePermissions.Ctx(util.Ctx).
		Where("p_type", "p").
		WhereIn("v0", roleIds).
		Where("v1", "backend_api").
		Fields("v2").
		Array()
	if err != nil {
		panic(err)
	}
	return myPermissions
}
