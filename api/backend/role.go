package backend

import "github.com/gogf/gf/v2/frame/g"

type RoleIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"角色列表"`
	Page   int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size   int `json:"size" in:"query" d:"15" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}
type RoleIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
