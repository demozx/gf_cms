package backend

import "github.com/gogf/gf/v2/frame/g"

type AdListIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台广告列表"`
}

type AdListIndexRes struct{}

type AdListAddReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台添加广告"`
}
type AdListAddRes struct{}

type AdListEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台编辑广告"`
	Id     int `json:"id" dc:"广告id"`
}
type AdListEditRes struct{}
