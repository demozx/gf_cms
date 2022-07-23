package model

import "github.com/gogf/gf/v2/util/gmeta"

type RoleItem struct {
	gmeta.Meta  `orm:"table:role"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	IsEnable    int    `json:"is_enable"`
	IsSystem    int    `json:"is_system"`
}

type RoleTitle struct {
	gmeta.Meta `orm:"table:role"`
	Id         int    `json:"id"`
	Title      string `json:"title"`
}
