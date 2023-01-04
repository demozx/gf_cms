package pc

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta    `tags:"Pc" method:"get" summary:"pc文章列表"`
	ChannelId int `json:"id" dc:"文章栏目id" d:"1"`
	model.PageSizeReq
}
type ArticleListRes struct {
	List  []*model.ArticleListItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type ArticleDetailReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"pc文章详情"`
	Id     int `json:"id" dc:"文章id"`
}
type ArticleDetailRes struct{}
