package backendApi

import "github.com/gogf/gf/v2/frame/g"

type FriendlyLinkStatusReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"友情链接修改状态"`
	Id     int `json:"id" dc:"友情链接id"`
}
type FriendlyLinkStatusRes struct{}

type FriendlyLinkAddReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"友情链接修改状态"`
	Name   string `json:"name" dc:"链接名称" v:"required#链接名称必填"`
	Url    string `json:"url" dc:"链接地址" v:"required|url#链接地址必填|链接地址格式错误"`
}
type FriendlyLinkAddRes struct{}
