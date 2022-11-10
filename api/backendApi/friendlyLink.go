package backendApi

import "github.com/gogf/gf/v2/frame/g"

type FriendlyLinkStatusReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"友情链接修改状态"`
	Id     int `json:"id" dc:"友情链接id"`
}
type FriendlyLinkStatusRes struct{}
