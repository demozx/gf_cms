// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"gf_cms/internal/model"
)

type (
	IMenu interface {
		// BackendView 获取全部后台菜单
		BackendView() (backendView []model.MenuGroups, err error)
		// BackendApi 获取全部后台菜单接口
		BackendApi() (res []model.MenuGroups, err error)
		// BackendMyMenu 我的后台菜单
		BackendMyMenu(accountId string) (backendMyMenus []model.MenuGroups, err error)
		BackendMyApi(accountId string) (backendMyMenus []model.MenuGroups, err error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
