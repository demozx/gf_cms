package backend

import "github.com/gogf/gf/v2/frame/g"

type ChannelIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"栏目分类列表"`
}
type ChannelIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ChannelAddReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"栏目分类添加"`
	Id     int `json:"id" in:"query" d:"0"  v:"min:0#id错误"     dc:"id，默认0"`
}
type ChannelAddRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ChannelEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"栏目分类编辑"`
	Id     int `json:"id" in:"query" d:"0"  v:"min:0#id错误"     dc:"id，默认0"`
}
type ChannelEditRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type ChannelModelIndexReq struct {
	g.Meta    `tags:"Backend" method:"get" summary:"数据列表"`
	Type      string `json:"type"`
	ChannelId int    `json:"channel_id"`
	WithTab   bool   `json:"with_tab" d:"true" v:"in:true,false#with_tab不合法"`
}
type ChannelModelIndexRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
