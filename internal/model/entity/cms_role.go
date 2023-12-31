// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CmsRole is the golang structure for table cms_role.
type CmsRole struct {
	Id          uint64 `json:"id"          ` //
	Title       string `json:"title"       ` // 中文名称
	Description string `json:"description" ` //
	Type        string `json:"type"        ` // 账户类型
	IsEnable    int    `json:"isEnable"    ` // 是否可用
	IsSystem    int    `json:"isSystem"    ` //
}
