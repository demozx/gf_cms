// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CmsAdmin is the golang structure of table cms_admin for DAO operations like Where/Data.
type CmsAdmin struct {
	g.Meta    `orm:"table:cms_admin, do:true"`
	Id        interface{} //
	Username  interface{} // 用户名
	Password  interface{} // 密码
	Name      interface{} // 姓名
	Tel       interface{} // 手机号
	Email     interface{} // 邮箱
	Status    interface{} // 状态
	IsSystem  interface{} // 是否系统用户
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
