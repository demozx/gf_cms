package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	insMenu = sMenu{}
)

//菜单
type sMenu struct{}

func Menu() *sMenu {
	return &insMenu
}

type MenuGroups struct {
	Title    string         `yaml:"title"`
	Children []MenuChildren `yaml:"children"`
}
type MenuChildren struct {
	Title      string `yaml:"title"`
	Route      string `yaml:"route"`
	Permission string `yaml:"permission"`
}
type Menus struct {
	Title  string       `yaml:"title"`
	Groups []MenuGroups `yaml:"groups"`
}
type MenuConfig struct {
	Backend Menus `yaml:"backend"`
	Web     Menus `yaml:"web"`
}

func (*sMenu) readYamlConfig(path string) (*MenuConfig, error) {
	conf := &MenuConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	//fmt.Println("conf: ", conf)
	return conf, nil
}

func (*sMenu) readYaml() *MenuConfig {
	var SystemRoot = Util().GetConfig("server.systemRoot")
	conf, err := Menu().readYamlConfig(SystemRoot + "/manifest/config/menu.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendAll Backend 获取全部后台菜单
func (*sMenu) BackendAll() []MenuGroups {
	cacheKey := PublicCachePreFix + ":menus:backend_all"
	result, err := g.Redis().Do(Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var menuGroups []MenuGroups
		if err = result.Structs(&menuGroups); err != nil {
			panic(err)
		}
		return menuGroups
	}
	backendAll := Menu().readYaml().Backend.Groups
	_, err = g.Redis().Do(Ctx, "SET", cacheKey, backendAll)
	if err != nil {
		panic(err)
	}
	return backendAll
}

// BackendMy 我的后台菜单
func (*sMenu) BackendMy(accountId string) []MenuGroups {
	//accountId := Middleware().GetAdminUserID(r)
	backendMyPermissions := Permission().BackendMy(accountId)
	backendAllMenus := Menu().BackendAll()
	var backendMyMenus []MenuGroups
	var backendMyMenusChildren []MenuChildren
	for _, menu := range backendAllMenus {
		var title = menu.Title
		var children = menu.Children
		for _, item := range children {
			var childrenPermission = item.Permission
			for _, myPermission := range backendMyPermissions {
				if myPermission.String() == childrenPermission {
					backendMyMenusChildren = append(backendMyMenusChildren, item)
				}
			}
		}
		var backendMyMenu MenuGroups
		backendMyMenu.Title = title
		backendMyMenu.Children = backendMyMenusChildren
		backendMyMenus = append(backendMyMenus, backendMyMenu)
	}
	//g.Log().Info(Ctx, "backendAllMenus", backendAllMenus)
	//g.Log().Info(Ctx, "backendMyMenus", backendMyMenus)
	//g.Log().Info(Ctx, "backendMyMenusChildren", backendMyMenusChildren)
	return backendMyMenus
}
