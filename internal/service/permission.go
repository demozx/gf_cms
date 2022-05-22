package service

import (
	"fmt"
	"gf_cms/internal/service/internal/dao"
	"github.com/gogf/gf/v2/database/gdb"
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
	backendAll := Permission().readYaml().Backend.Groups
	//fmt.Println(backendAll)
	return backendAll
}

// BackendMy 获取我的所有后台权限
func (*sPermission) BackendMy(accountId int) []gdb.Value {
	roleIds := GetRoleIdsByAccountId(accountId)
	myPermissions, err := dao.CmsRulePermissions.Ctx(Ctx).
		Where("p_type", "p").
		WhereIn("v0", roleIds).
		Where("v1", "backend").
		Fields("v2").
		Array()
	if err != nil {
		panic(err)
	}
	fmt.Println("myPermissions", myPermissions)
	return myPermissions
}
