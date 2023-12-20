// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsShortcut is the golang structure for table cms_shortcut.
type CmsShortcut struct {
	Id        uint        `json:"id"        ` //
	AccountId uint        `json:"accountId" ` // 用户id
	Name      string      `json:"name"      ` // 快捷方式名称
	Route     string      `json:"route"     ` // 路由
	Sort      int         `json:"sort"      ` // 排序
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
