package service

import (
	"gf_cms/internal/service/internal/dao"
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

// Permission 权限
func Permission() *sPermission {
	return &insPermission
}

type PermissionGroups struct {
	Slug        string                  `yaml:"slug"`
	Title       string                  `yaml:"title"`
	Permissions []PermissionPermissions `yaml:"permissions"`
}
type PermissionPermissions struct {
	Title string `yaml:"title"`
	Slug  string `yaml:"slug"`
}
type Permissions struct {
	Title  string             `yaml:"title"`
	Groups []PermissionGroups `yaml:"groups"`
}
type PermissionConfig struct {
	Backend Permissions `yaml:"backend"`
	Web     Permissions `yaml:"web"`
}

func (*sPermission) readYamlConfig(path string) (*PermissionConfig, error) {
	conf := &PermissionConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	//fmt.Println("conf: ", conf)
	return conf, nil
}

func (*sPermission) readYaml() *PermissionConfig {
	var SystemRoot = Util().GetConfig("server.systemRoot")
	conf, err := Permission().readYamlConfig(SystemRoot + "/manifest/config/permission.yaml")
	//fmt.Println(conf)
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendAll Backend 获取全部后台权限
func (*sPermission) BackendAll() []PermissionGroups {
	cacheKey := PublicCachePreFix + ":permissions:backend_all"
	result, err := g.Redis().Do(Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var permissionGroups []PermissionGroups
		if err = result.Structs(&permissionGroups); err != nil {
			panic(err)
		}
		return permissionGroups
	}
	backendAll := Permission().readYaml().Backend.Groups
	_, err = g.Redis().Do(Ctx, "SET", cacheKey, backendAll)
	if err != nil {
		panic(err)
	}
	return backendAll
}

// BackendMy 获取我的所有后台权限
func (*sPermission) BackendMy(accountId string) []gdb.Value {
	roleIds := GetRoleIdsByAccountId(accountId)
	if len(roleIds) == 0 {
		panic("用户无任何角色")
	}
	myPermissions, err := dao.CmsRulePermissions.Ctx(Ctx).
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
