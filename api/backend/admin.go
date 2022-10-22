package backend

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type AdminLoginReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台登录"`
}
type AdminLoginRes struct{}

type AdminIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"管理员列表"`
	model.PageSizeReq
}
type AdminIndexRes struct{}
type AdminAddReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"添加管理员"`
}
type AdminAddRes struct{}

type AdminEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"编辑管理员"`
	Id     int `json:"id" in:"query" d:"0"  v:"min:0#id错误"     dc:"id，默认0"`
}
type AdminEditRes struct{}
