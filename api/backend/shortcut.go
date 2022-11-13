package backend

import "github.com/gogf/gf/v2/frame/g"

type ShortcutIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"快捷方式列表"`
}
type ShortcutIndexRes struct{}

type ShortcutEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"编辑快捷方式"`
	Id     int `json:"id" v:"required|min:1#快捷方式id必填|id错误"`
}
type ShortcutEditRes struct{}
