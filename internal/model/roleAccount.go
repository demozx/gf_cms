package model

import "github.com/gogf/gf/v2/util/gmeta"

type RoleAccountItem struct {
	gmeta.Meta `orm:"table:role_account"`
	Id         int `json:"id"`         //id
	AccountId  int `json:"account_id"` //账号id
	RoleId     int `json:"role_id"`    //角色id
	RoleTitle  `orm:"with:id=role_id"`
}
