package mobile

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type RouterBeautifyDefaultReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"mobile路由美化"`
}
type RouterBeautifyDefaultRes struct{}

type RouterBeautifyPageSizeReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"mobile路由美化pageSizeReq"`
	model.PageSizeReq
}

type RouterBeautifyDetailReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"mobile路由美化detailReq"`
	Id     int `json:"id" dc:"详情id"`
}
