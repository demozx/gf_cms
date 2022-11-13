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

type FriendlyLinkEditReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"友情链接修改状态"`
	Id     int    `json:"id" d:"0" v:"required|min:1#id必填|id不合法"`
	Name   string `json:"name" dc:"链接名称" v:"required#链接名称必填"`
	Url    string `json:"url" dc:"链接地址" v:"required|url#链接地址必填|链接地址格式错误"`
}
type FriendlyLinkEditRes struct{}

type FriendlyLinkSortReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"友情链接排序"`
	Sort   []string `json:"sort" dc:"排序" arg:"true" v:"required#排序字段必填"`
}
type FriendlyLinkSortRes struct{}

type FriendlyLinkBatchDeleteReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"友情链接批量删除"`
	Ids    []int `json:"ids" v:"required#id必填"`
}
type FriendlyLinkBatchDeleteRes struct{}
