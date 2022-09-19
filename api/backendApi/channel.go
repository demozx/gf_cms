package backendApi

import "github.com/gogf/gf/v2/frame/g"

type ChannelIndexApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目分类列表"`
}
type ChannelIndexApiRes struct {
}

type ChannelStatusApiReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"栏目启用禁用"`
	Id     int `p:"id" name:"id" brief:"栏目ID" des:"栏目ID"  arg:"true" v:"required#请选择栏目ID"`
}
type ChannelStatusApiRes struct {
}
