package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta    `tags:"BackendApi" method:"post" summary:"文章列表"`
	ChannelId int    `p:"channel_id" name:"channel_id" v:"" des:"频道ID" brief:"频道ID" arg:"true"`
	StartAt   string `p:"start_at" name:"start_at" v:"date-format:Y-m-d H:i:s" des:"开始时间" brief:"频道ID" arg:"true"`
	EndAt     string `p:"end_at" name:"end_at" v:"date-format:Y-m-d H:i:s" des:"开始时间" brief:"频道ID" arg:"true"`
	Keyword   string `p:"keyword" name:"keyword" v:"" des:"关键词"`
	model.PageSizeReq
}
type ArticleListRes struct {
	List  []model.ArticleListItem `p:"list" name:"list" des:"列表"`
	Page  int                     `p:"page" name:"page" des:"分页码"`
	Size  int                     `p:"size" name:"size" des:"分页数量"`
	Total int                     `p:"total" name:"total" des:"数据总数"`
}

type ArticleSortReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"文章排序"`
	Sort   []string `p:"sort" name:"sort" des:"排序"`
}
type ArticleSortRes struct{}

type ArticleFlagReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"flag"`
	Ids    []int  `p:"ids" name:"ids" dsc:"文章ID们" v:"required#文章ID必填" arg:"true"`
	Flag   string `p:"flag" name:"flag" dec:"flag" v:"in:t,r#类型不合法" arg:"true"`
}
type ArticleFlagRes struct{}

type ArticleStatusReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"审核状态"`
	Ids    []int `p:"ids" name:"ids" dsc:"文章ID们" v:"required#文章ID必填" arg:"true"`
}
type ArticleStatusRes struct{}
