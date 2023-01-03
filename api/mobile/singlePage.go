package mobile

import "github.com/gogf/gf/v2/frame/g"

type SinglePageReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"移动单页"`
	Id     int `json:"id" dc:"栏目id" d:"1"`
}
type SinglePageRes struct{}
