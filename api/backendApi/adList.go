package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type AdListIndexReq struct {
	g.Meta    `tags:"Backend" method:"post" summary:"后台广告列表"`
	ChannelId int `json:"channel_id" description:"分类id"`
	model.PageSizeReq
}

type AdListIndexRes struct {
	List  []*model.AdListItem `json:"list" description:"后台广告列表接口结果"`
	Page  int                 `json:"page" description:"分页码"`
	Size  int                 `json:"size" description:"分页数量"`
	Total int                 `json:"total" description:"数据总数"`
}
