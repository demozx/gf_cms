package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ImageListReq struct {
	g.Meta    `tags:"BackendApi" method:"post" summary:"图集列表"`
	ChannelId int    `name:"channel_id" v:"" dc:"频道ID" brief:"频道ID" arg:"true"`
	StartAt   string `name:"start_at" v:"date-format:Y-m-d H:i:s" dc:"开始时间" brief:"频道ID" arg:"true"`
	EndAt     string `name:"end_at" v:"date-format:Y-m-d H:i:s" dc:"开始时间" brief:"频道ID" arg:"true"`
	Keyword   string `name:"keyword" v:"" dc:"关键词"`
	model.PageSizeReq
}
type ImageListRes struct {
	List  []model.ArticleListItem `name:"list" dc:"图集列表"`
	Page  int                     `name:"page" dc:"分页码"`
	Size  int                     `name:"size" dc:"分页数量"`
	Total int                     `name:"total" dc:"数据总数"`
}

type ImageSortReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"图集排序"`
	Sort   []string `name:"sort" dc:"排序"`
}
type ImageSortRes struct{}

type ImageFlagReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"图集flag"`
	Ids    []int  `name:"ids" dc:"文章ID们" v:"required#图集ID必填" arg:"true"`
	Flag   string `name:"flag" dc:"flag" v:"in:t,r#类型不合法" arg:"true"`
}
type ImageFlagRes struct{}

type ImageStatusReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"开启关闭"`
	Ids    []int `name:"ids" dc:"图集ID们" v:"required#图集ID必填" arg:"true"`
}
type ImageStatusRes struct{}

type ImageDeleteReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"删除图集"`
	Ids    []int `name:"ids" dc:"图集ID们" v:"required#图集ID必填" arg:"true"`
}
type ImageDeleteRes struct{}

type ImageMoveReq struct {
	g.Meta    `tags:"BackendApi" method:"post" summary:"移动图集"`
	ChannelId int    `name:"channel_id" dc:"频道ID" v:"required#频道ID必填" arg:"true"`
	StrIds    string `name:"str_ids" dc:"ids，英文逗号拼接" v:"required#必填项不能为空" arg:"true"`
}
type ImageMoveRes struct{}
