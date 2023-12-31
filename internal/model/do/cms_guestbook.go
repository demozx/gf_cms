// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsGuestbook is the golang structure of table cms_guestbook for DAO operations like Where/Data.
type CmsGuestbook struct {
	g.Meta    `orm:"table:cms_guestbook, do:true"`
	Id        interface{} //
	Name      interface{} // 留言者姓名
	Tel       interface{} // 留言者电话
	Content   interface{} // 留言内容
	From      interface{} // 来源：1、电脑端，2、移动端
	Ip        interface{} // 留言者ip
	Address   interface{} // 留言者归属地
	Status    interface{} // 留言状态：0、未读，1、已读
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
