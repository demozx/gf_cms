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

type ImageAddReq struct {
	g.Meta      `tags:"BackendApi" method:"post" summary:"新增图集"`
	ChannelId   int    `name:"channel_id" dc:"频道ID" v:"required#频道ID必填" arg:"true"`
	Title       string `name:"title" dc:"图集标题" v:"required#图集标题必填" arg:"true"`
	Images      string `name:"images" dc:"图片们" v:"required#图片必传" arg:"true"`
	Description string `name:"description" dc:"文章摘要" v:"" arg:"true"`
	FlagR       int    `name:"flag_r" dc:"文章属性-推荐" v:"in:0,1#推荐属性不合法" arg:"true"`
	FlagT       int    `name:"flag_t" dc:"文章属性-置顶" v:"in:0,1#置顶属性不合法" arg:"true"`
	Status      int    `name:"status" dc:"审核状态" v:"in:0,1#审核状态不合法" arg:"true"`
	CreatedAt   string `name:"created_at" dc:"发布时间" v:"date-format:Y-m-d H:i:s" arg:"true"`
	ClickNum    int    `name:"copy_from" dc:"点击数" v:"" arg:"true"`
}
type ImageAddRes struct{}

type ImageEditReq struct {
	g.Meta      `tags:"BackendApi" method:"post" summary:"新增图集"`
	Id          int    `name:"id" dc:"图集id" v:"required#图集id必填" arg:"true"`
	ChannelId   int    `name:"channel_id" dc:"频道ID" v:"required#频道ID必填" arg:"true"`
	Title       string `name:"title" dc:"图集标题" v:"required#图集标题必填" arg:"true"`
	Images      string `name:"images" dc:"图片们" v:"required#图片必传" arg:"true"`
	Description string `name:"description" dc:"文章摘要" v:"" arg:"true"`
	FlagR       int    `name:"flag_r" dc:"文章属性-推荐" v:"in:0,1#推荐属性不合法" arg:"true"`
	FlagT       int    `name:"flag_t" dc:"文章属性-置顶" v:"in:0,1#置顶属性不合法" arg:"true"`
	Status      int    `name:"status" dc:"审核状态" v:"in:0,1#审核状态不合法" arg:"true"`
	CreatedAt   string `name:"created_at" dc:"发布时间" v:"date-format:Y-m-d H:i:s" arg:"true"`
	ClickNum    int    `name:"copy_from" dc:"点击数" v:"" arg:"true"`
}
type ImageEditRes struct{}

type ImageBatchDestroyReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"回收站-批量删除图集"`
	Ids    []int `name:"ids" dc:"图集ID们" v:"required#图集ID必填" arg:"true"`
}
type ImageBatchDestroyRes struct{}

type ImageBatchRestoreReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"回收站-批量恢复图集"`
	Ids    []int `name:"ids" dc:"图集ID们" v:"required#图集ID必填" arg:"true"`
}
type ImageBatchRestoreRes struct{}
