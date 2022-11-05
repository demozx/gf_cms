package backend

import "github.com/gogf/gf/v2/frame/g"

type AdChannelIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台广告分类列表"`
}

type AdChannelIndexRes struct{}

type AdChannelEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台广告分类编辑"`
	Id     int `json:"id" in:"query" d:"0"  v:"min:0#id错误"     dc:"id，默认0"`
}

type AdChannelEditRes struct{}
