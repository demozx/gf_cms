package permission

import (
	"gf_cms/internal/dao"
	"gf_cms/internal/logic/admin"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
	"log"
	"os"
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

// BackendAll Backend 获取全部后台权限
func (*sPermission) BackendAll() []model.PermissionGroups {
	cacheKey := util.PublicCachePreFix + ":permissions:backend_all"
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
	backendAll := Permission().readYaml().Backend.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendAll)
	if err != nil {
		panic(err)
	}
	return backendAll
}

// BackendMy 获取我的所有后台权限
func (*sPermission) BackendMy(accountId string) []gdb.Value {
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
