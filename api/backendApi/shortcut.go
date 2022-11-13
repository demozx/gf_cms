package backendApi

import "github.com/gogf/gf/v2/frame/g"

type ShortcutAddReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"添加快捷方式"`
	Name   string `json:"name" v:"required#快捷方式名称必填"`
	Route  string `json:"route" v:"required#快捷方式路由必填"`
}
type ShortcutAddRes struct{}

type ShortcutEditReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"添加快捷方式"`
	Id     int    `json:"id" v:"required|min:1#快捷方式id必填|id错误"`
	Name   string `json:"name" v:"required#快捷方式名称必填"`
	Route  string `json:"route" v:"required#快捷方式路由必填"`
}
type ShortcutEditRes struct{}

type ShortcutBatchDeleteReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"批量删除快捷方式"`
	Ids    []int `json:"ids" v:"required#快捷方式ids必填"`
}
type ShortcutBatchDeleteRes struct{}
