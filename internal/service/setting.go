// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"gf_cms/internal/model"
)

type (
	ISetting interface {
		// BackendViewAll 获取所有后台菜单
		BackendViewAll() (backendAll []model.SettingGroups, err error)
		// Save 保存设置
		Save(forms map[string]interface{}) (res bool, err error)
	}
)

var (
	localSetting ISetting
)

func Setting() ISetting {
	if localSetting == nil {
		panic("implement not found for interface ISetting, forgot register?")
	}
	return localSetting
}

func RegisterSetting(i ISetting) {
	localSetting = i
}
