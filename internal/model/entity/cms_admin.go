// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsAdmin is the golang structure for table cms_admin.
type CmsAdmin struct {
	Id        uint        `json:"id"        ` //
	Username  string      `json:"username"  ` // 用户名
	Password  string      `json:"password"  ` // 密码
	Name      string      `json:"name"      ` // 姓名
	Tel       string      `json:"tel"       ` // 手机号
	Email     string      `json:"email"     ` // 邮箱
	Status    int         `json:"status"    ` // 状态
	IsSystem  int         `json:"isSystem"  ` // 是否系统用户
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
