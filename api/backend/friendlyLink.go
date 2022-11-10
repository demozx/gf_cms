package backend

import "github.com/gogf/gf/v2/frame/g"

type FriendlyLinkIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"友情链接列表"`
}
type FriendlyLinkIndexRes struct{}
