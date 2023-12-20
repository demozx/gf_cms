// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsShortcut is the golang structure of table cms_shortcut for DAO operations like Where/Data.
type CmsShortcut struct {
	g.Meta    `orm:"table:cms_shortcut, do:true"`
	Id        interface{} //
	AccountId interface{} // 用户id
	Name      interface{} // 快捷方式名称
	Route     interface{} // 路由
	Sort      interface{} // 排序
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
