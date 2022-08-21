package model

import "github.com/gogf/gf/v2/util/gmeta"

type RoleItem struct {
	gmeta.Meta  `orm:"table:role"`
	Id          int                   `json:"id"`
	Name        string                `json:"name"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Type        string                `json:"type"`
	IsEnable    int                   `json:"is_enable"`
	IsSystem    int                   `json:"is_system"`
	Permissions []RulePermissionsItem `orm:"with:v0=id"`
}

type RoleTitle struct {
	gmeta.Meta `orm:"table:role"`
	Id         int    `json:"id"`
	Title      string `json:"title"`
}

type RoleGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
}
type RoleListItem struct {
	gmeta.Meta  `orm:"table:role"`
	Id          int                   `json:"id"`
	Name        string                `json:"name"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Type        string                `json:"type"`
	IsEnable    int                   `json:"is_enable"`
	IsSystem    int                   `json:"is_system"`
	Permissions []RulePermissionsItem `orm:"with:v0=id"`
}
type RoleGetListOutput struct {
	List  []RoleListItem `json:"list" description:"列表"`
	Page  int            `json:"page" description:"分页码"`
	Size  int            `json:"size" description:"分页数量"`
	Total int            `json:"total" description:"数据总数"`
}
