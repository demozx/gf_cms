// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CmsRoleAccount is the golang structure of table cms_role_account for DAO operations like Where/Data.
type CmsRoleAccount struct {
	g.Meta    `orm:"table:cms_role_account, do:true"`
	Id        interface{} // ID
	AccountId interface{} // 账户id
	RoleId    interface{} // 角色id
}
