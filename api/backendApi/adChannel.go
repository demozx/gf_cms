package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type AdChannelIndexReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"后台广告分类列表接口参数"`
	model.PageSizeReq
}

type AdChannelIndexRes struct {
	List  []*model.AdChannelListItem `json:"list" description:"后台广告分类列表接口结果"`
	Page  int                        `json:"page" description:"分页码"`
	Size  int                        `json:"size" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type AdChannelAddReq struct {
	g.Meta      `tags:"Backend" method:"post" summary:"新增广告分类"`
	ChannelName string `json:"channel_name" description:"分类名称"`
	Remarks     string `json:"remarks" description:"备注"`
}
type AdChannelAddRes struct{}

type AdChannelEditReq struct {
	g.Meta    `tags:"Backend" method:"post" summary:"编辑广告分类"`
	ChannelId int `json:"channel_id" description:"分类id"`
	AdChannelAddReq
}
type AdChannelEditRes struct{}

type AdChannelDeleteReq struct {
	g.Meta    `tags:"Backend" method:"post" summary:"删除广告分类"`
	ChannelId int `json:"channel_id" description:"分类id"`
}
type AdChannelDeleteRes struct{}

type AdChannelSortReq struct {
	g.Meta `tags:"Backend" method:"post" summary:"广告分类排序"`
	Sort   []string `name:"sort" dc:"排序"`
}
type AdChannelSortRes struct{}
