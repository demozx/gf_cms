package backend

import "github.com/gogf/gf/v2/frame/g"

type FriendlyLinkIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"友情链接列表"`
}
type FriendlyLinkIndexRes struct{}

type FriendlyLinkEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"编辑友情链接"`
	Id     int `json:"id" d:"0" v:"required|min:1#id必填|id不合法"`
}
type FriendlyLinkEditRes struct{}
