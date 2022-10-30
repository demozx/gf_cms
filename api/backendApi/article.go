package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta    `tags:"BackendApi" method:"post" summary:"文章列表"`
	ChannelId int    `name:"channel_id" v:"" dc:"频道ID" brief:"频道ID" arg:"true"`
	StartAt   string `name:"start_at" v:"date-format:Y-m-d H:i:s" dc:"开始时间" brief:"频道ID" arg:"true"`
	EndAt     string `name:"end_at" v:"date-format:Y-m-d H:i:s" dc:"开始时间" brief:"频道ID" arg:"true"`
	Keyword   string `name:"keyword" v:"" dc:"关键词"`
	model.PageSizeReq
}
type ArticleListRes struct {
	List  []model.ArticleListItem `name:"list" dc:"列表"`
	Page  int                     `name:"page" dc:"分页码"`
	Size  int                     `name:"size" dc:"分页数量"`
	Total int                     `name:"total" dc:"数据总数"`
}

type ArticleSortReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"文章排序"`
	Sort   []string `name:"sort" dc:"排序"`
}
type ArticleSortRes struct{}

type ArticleFlagReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"flag"`
	Ids    []int  `name:"ids" dc:"文章ID们" v:"required#文章ID必填" arg:"true"`
	Flag   string `name:"flag" dc:"flag" v:"in:t,r#类型不合法" arg:"true"`
}
type ArticleFlagRes struct{}

type ArticleStatusReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"审核状态"`
	Ids    []int `name:"ids" dc:"文章ID们" v:"required#文章ID必填" arg:"true"`
}
type ArticleStatusRes struct{}

type ArticleDeleteReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"删除"`
	Ids    []int `name:"ids" dc:"文章ID们" v:"required#文章ID必填" arg:"true"`
}
type ArticleDeleteRes struct{}

type ArticleMoveReq struct {
	g.Meta    `tags:"BackendApi" method:"post" summary:"移动"`
	ChannelId int    `name:"channel_id" dc:"频道ID" v:"required#频道ID必填" arg:"true"`
	StrIds    string `name:"str_ids" dc:"ids，英文逗号拼接" v:"required#必填项不能为空" arg:"true"`
}
type ArticleMoveRes struct{}

type ArticleAddReq struct {
	g.Meta      `tags:"BackendApi" method:"post" summary:"移动"`
	Title       string `name:"title" dc:"文章标题" v:"required#文章标题必填" arg:"true"`
	ChannelId   int    `name:"channel_id" dc:"频道ID" v:"required#频道ID必填" arg:"true"`
	Keyword     string `name:"keyword" dc:"关键词" v:"" arg:"true"`
	Description string `name:"description" dc:"文章摘要" v:"" arg:"true"`
	FlagP       int    `name:"flag_p" dc:"文章属性-带图" v:"in:0,1#带图属性不合法" arg:"true"`
	FlagR       int    `name:"flag_r" dc:"文章属性-推荐" v:"in:0,1#推荐属性不合法" arg:"true"`
	FlagT       int    `name:"flag_t" dc:"文章属性-置顶" v:"in:0,1#置顶属性不合法" arg:"true"`
	Status      int    `name:"status" dc:"审核状态" v:"in:0,1#审核状态不合法" arg:"true"`
	Thumb       string `name:"thumb" dc:"文章缩略图" v:"" arg:"true"`
	CreatedAt   string `name:"created_at" dc:"发布时间" v:"date-format:Y-m-d H:i:s" arg:"true"`
	CopyFrom    string `name:"copy_from" dc:"文章来源" v:"" arg:"true"`
	ClickNum    int    `name:"copy_from" dc:"点击数" v:"" arg:"true"`
	Body        string `name:"body" dc:"文章内容" v:"" arg:"true"`
}
type ArticleAddRes struct{}
