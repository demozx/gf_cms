package backend

import "github.com/gogf/gf/v2/frame/g"

type ChannelIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"栏目分类列表"`
}
type ChannelIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
