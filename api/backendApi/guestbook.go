package backendApi

import "github.com/gogf/gf/v2/frame/g"

type GuestbookStatusReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"修改留言状态"`
	Id     int `json:"id" dc:"留言id"`
}
type GuestbookStatusRes struct{}

type GuestbookDeleteReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"删除留言"`
	Ids    []int `json:"ids" dc:"留言ids"`
}
type GuestbookDeleteRes struct{}
