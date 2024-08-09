package menu

import (
	"context"
	"gf_cms/internal/logic/permission"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	insMenu = sMenu{}
)

// 菜单
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

func (*sMenu) readYaml(ctx context.Context) (conf *model.MenuConfig, err error) {
	data, err := g.Cfg("menu").Data(ctx)
	if err != nil {
		return nil, err
	}
	err = gconv.Scan(data, &conf)
	if err != nil {
		return nil, err
	}
	return
}

// BackendView 获取全部后台菜单
func (*sMenu) BackendView() (backendView []model.MenuGroups, err error) {
	cacheKey := util.PublicCachePreFix + ":menus:backend_view"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		return nil, err
	}
	if !result.IsEmpty() {
		var menuGroups []model.MenuGroups
		if err = result.Structs(&menuGroups); err != nil {
			return nil, err
		}
		return menuGroups, nil
	}
	conf, _ := Menu().readYaml(util.Ctx)
	backendView = conf.Backend.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendView)
	if err != nil {
		return nil, err
	}
	return backendView, nil
}

// BackendApi 获取全部后台菜单接口
func (*sMenu) BackendApi() (res []model.MenuGroups, err error) {
	cacheKey := util.PublicCachePreFix + ":menus:backend_api"
	result, err := g.Redis().Do(util.Ctx, "GET", cacheKey)
	if err != nil {
		return nil, err
	}
	if !result.IsEmpty() {
		var menuGroups []model.MenuGroups
		if err = result.Structs(&menuGroups); err != nil {
			return nil, err
		}
		return menuGroups, nil
	}
	conf, _ := Menu().readYaml(util.Ctx)
	backendApi := conf.BackendApi.Groups
	_, err = g.Redis().Do(util.Ctx, "SET", cacheKey, backendApi)
	if err != nil {
		return nil, err
	}
	return backendApi, nil
}

// BackendMyMenu 我的后台菜单
func (*sMenu) BackendMyMenu(accountId string) (backendMyMenus []model.MenuGroups, err error) {
	//accountId := Middleware().GetAdminUserID(r)
	backendMyPermissions, err := permission.Permission().BackendMyView(accountId)
	if err != nil {
		return nil, err
	}
	//g.Log().Info(Ctx, "backendMyPermissions", backendMyPermissions)
	backendViewMenus, err := Menu().BackendView()
	if err != nil {
		return nil, err
	}
	//var backendMyMenus []model.MenuGroups

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
	return backendMyMenus, err
}

func (*sMenu) BackendMyApi(accountId string) (backendMyMenus []model.MenuGroups, err error) {
	//accountId := Middleware().GetAdminUserID(r)
	backendMyPermissions, err := permission.Permission().BackendMyApi(accountId)
	if err != nil {
		return nil, err
	}
	backendApiMenus, err := Menu().BackendApi()
	if err != nil {
		return nil, err
	}
	//var backendMyMenus []model.MenuGroups
	for _, menu := range backendApiMenus {
		var title = menu.Title
		var children = menu.Children
		var backendMyMenusChildren []model.MenuChildren
		for _, item := range children {
			var childrenPermission = item.Permission
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
	return backendMyMenus, err
}
