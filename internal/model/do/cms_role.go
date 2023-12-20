// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CmsRole is the golang structure of table cms_role for DAO operations like Where/Data.
type CmsRole struct {
	g.Meta      `orm:"table:cms_role, do:true"`
	Id          interface{} //
	Title       interface{} // 中文名称
	Description interface{} //
	Type        interface{} // 账户类型
	IsEnable    interface{} // 是否可用
	IsSystem    interface{} //
}
