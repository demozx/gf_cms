// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsFriendlyLink is the golang structure for table cms_friendly_link.
type CmsFriendlyLink struct {
	Id        uint        `json:"id"        ` //
	Name      string      `json:"name"      ` // 链接名称
	Url       string      `json:"url"       ` // 链接地址
	Status    int         `json:"status"    ` // 状态
	Sort      int         `json:"sort"      ` // 排序
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
