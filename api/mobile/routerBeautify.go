package mobile

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RouterBeautifyDefaultReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"mobile路由美化"`
}
type RouterBeautifyDefaultRes struct{}

type RouterBeautifyPageSizeReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"mobile路由美化pageSizeReq"`
	Page   int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size   int `json:"size" in:"query" d:"10" v:"max:100#分页数量最大100条" dc:"分页数量，最大100"`
}

type RouterBeautifyDetailReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"mobile路由美化detailReq"`
	Id     int `json:"id" dc:"详情id"`
}
