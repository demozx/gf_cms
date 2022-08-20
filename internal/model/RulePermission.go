package model

import "github.com/gogf/gf/v2/util/gmeta"

type RulePermissionsItem struct {
	gmeta.Meta `orm:"table:rule_permissions"`
	PType      string `json:"p_type"`
	V0         string `json:"v0"`
	V1         string `json:"v1"`
	V2         string `json:"v2"`
	Title      string `json:"title"`
}
