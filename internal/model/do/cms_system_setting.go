// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsSystemSetting is the golang structure of table cms_system_setting for DAO operations like Where/Data.
type CmsSystemSetting struct {
	g.Meta    `orm:"table:cms_system_setting, do:true"`
	Id        interface{} //
	Group     interface{} //
	Name      interface{} // 名称
	Value     interface{} // 值
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
