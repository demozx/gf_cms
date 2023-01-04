package pc

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type SearchReq struct {
	g.Meta  `tags:"Pc" method:"get" summary:"pc搜索"`
	Keyword string `json:"keyword" dc:"关键词"`
	Type    string `json:"type" dc:"模型" d:"article"`
	model.PageSizeReq
}
type SearchRes struct {
	List  []*model.ArticleListItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}
