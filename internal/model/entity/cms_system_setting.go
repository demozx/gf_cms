// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsSystemSetting is the golang structure for table cms_system_setting.
type CmsSystemSetting struct {
	Id        uint        `json:"id"        ` //
	Group     string      `json:"group"     ` //
	Name      string      `json:"name"      ` // 名称
	Value     string      `json:"value"     ` // 值
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
