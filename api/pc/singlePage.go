package pc

import "github.com/gogf/gf/v2/frame/g"

type SinglePageReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"pc单页"`
	Id     int `json:"id" dc:"栏目id" d:"1"`
}
type SinglePageRes struct{}
