package backend

import "github.com/gogf/gf/v2/frame/g"

type RecycleBinIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"回收站列表"`
	Type   string `json:"type"`
	Page   int    `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size   int    `json:"size" in:"query" d:"15" v:"max:50#分页数量最大50条" dc:"分页数量，最大50"`
}
type RecycleBinIndexRes struct{}
