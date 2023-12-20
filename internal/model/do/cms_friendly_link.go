// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsFriendlyLink is the golang structure of table cms_friendly_link for DAO operations like Where/Data.
type CmsFriendlyLink struct {
	g.Meta    `orm:"table:cms_friendly_link, do:true"`
	Id        interface{} //
	Name      interface{} // 链接名称
	Url       interface{} // 链接地址
	Status    interface{} // 状态
	Sort      interface{} // 排序
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
