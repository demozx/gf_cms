package menu

import (
	"gf_cms/internal/logic/permission"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"log"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"gopkg.in/yaml.v3"
)

var (
	insMenu = sMenu{}
)

//菜单
type sMenu struct{}

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

func Menu() *sMenu {
	return &insMenu
}

func (*sMenu) readYamlConfig(path string) (*model.MenuConfig, error) {
	conf := &model.MenuConfig{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
	return conf, nil
}

func (*sMenu) readYaml() *model.MenuConfig {
	conf, err := Menu().readYamlConfig(util.Util().SystemRoot() + "/manifest/config/menu.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return conf
}

// BackendView 获取全部后台菜单
func (*sMenu) BackendView() []model.MenuGroups {
	cacheKey := util.PublicCachePreFix + ":menus:backend_view"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var menuGroups []model.MenuGroups
		if err = result.Structs(&menuGroups); err != nil {
			panic(err)
		}
		return menuGroups
	}
	backendView := Menu().readYaml().Backend.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendView)
	if err != nil {
		panic(err)
	}
	return backendView
}

// BackendApi 获取全部后台菜单接口
func (*sMenu) BackendApi() []model.MenuGroups {
	cacheKey := util.PublicCachePreFix + ":menus:backend_api"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		panic(err)
	}
	if !result.IsEmpty() {
		var menuGroups []model.MenuGroups
		if err = result.Structs(&menuGroups); err != nil {
			panic(err)
		}
		return menuGroups
	}
	backendApi := Menu().readYaml().BackendApi.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendApi)
	if err != nil {
		panic(err)
	}
	return backendApi
}

// BackendMyMenu 我的后台菜单
func (*sMenu) BackendMyMenu(accountId string) []model.MenuGroups {
	//accountId := Middleware().GetAdminUserID(r)
	backendMyPermissions := permission.Permission().BackendMyView(accountId)
	//g.Log().Info(Ctx, "backendMyPermissions", backendMyPermissions)
	backendViewMenus := Menu().BackendView()
	var backendMyMenus []model.MenuGroups

	for _, menu := range backendViewMenus {
		var title = menu.Title
		var children = menu.Children
		var backendMyMenusChildren []model.MenuChildren
		for _, item := range children {
			var childrenPermission = item.Permission
			//g.Log().Info(Ctx, "childrenPermission", childrenPermission)
			for _, myPermission := range backendMyPermissions {
				//g.Log().Info(Ctx, "myPermission.String()", myPermission.String(), childrenPermission, myPermission.String() == childrenPermission)
				if myPermission.String() == childrenPermission {
					backendMyMenusChildren = append(backendMyMenusChildren, item)
				}
			}
		}
		//g.Log().Info(Ctx, "backendMyMenusChildren", backendMyMenusChildren)
		if backendMyMenusChildren != nil {
			var backendMyMenu model.MenuGroups
			backendMyMenu.Title = title
			backendMyMenu.Children = backendMyMenusChildren
			backendMyMenus = append(backendMyMenus, backendMyMenu)
			//g.Log().Info(Ctx, "backendMyMenu", backendMyMenu)
		}
	}
	//g.Log().Info(Ctx, "backendAllMenus", backendAllMenus)
	//g.Log().Info(Ctx, "backendMyMenus", backendMyMenus)
	return backendMyMenus
}

func (*sMenu) BackendMyApi(accountId string) []model.MenuGroups {
	//accountId := Middleware().GetAdminUserID(r)
	backendMyPermissions := permission.Permission().BackendMyApi(accountId)
	//g.Log().Info(Ctx, "backendMyPermissions", backendMyPermissions)
	backendApiMenus := Menu().BackendApi()
	var backendMyMenus []model.MenuGroups

	for _, menu := range backendApiMenus {
		var title = menu.Title
		var children = menu.Children
		var backendMyMenusChildren []model.MenuChildren
		for _, item := range children {
			var childrenPermission = item.Permission
			//g.Log().Info(Ctx, "childrenPermission", childrenPermission)
			for _, myPermission := range backendMyPermissions {
				//g.Log().Info(Ctx, "myPermission.String()", myPermission.String(), childrenPermission, myPermission.String() == childrenPermission)
				if myPermission.String() == childrenPermission {
					backendMyMenusChildren = append(backendMyMenusChildren, item)
				}
			}
		}
		//g.Log().Info(Ctx, "backendMyMenusChildren", backendMyMenusChildren)
		if backendMyMenusChildren != nil {
			var backendMyMenu model.MenuGroups
			backendMyMenu.Title = title
			backendMyMenu.Children = backendMyMenusChildren
			backendMyMenus = append(backendMyMenus, backendMyMenu)
			//g.Log().Info(Ctx, "backendMyMenu", backendMyMenu)
		}
	}
	//g.Log().Info(Ctx, "backendAllMenus", backendAllMenus)
	//g.Log().Info(Ctx, "backendMyMenus", backendMyMenus)
	return backendMyMenus
}
