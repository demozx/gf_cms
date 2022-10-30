package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleMoveReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"文章移动"`
	StrIds string `json:"str_ids" in:"query" d:""  v:"required#必填项不能为空"  dc:"str_ids，英文逗号拼接"`
}
type ArticleMoveRes struct{}

type ArticleAddReq struct {
	ChannelId int `json:"channel_id" in:"query" d:"0"  v:""  dc:"频道ID"`
	g.Meta    `tags:"Backend" method:"get" summary:"文章新增"`
}
type ArticleAddRes struct{}

type ArticleEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"文章编辑"`
	Id     int `json:"id" in:"query" d:"0"  v:"required#文章ID必填"  dc:"文章ID"`
}
type ArticleEditRes struct{}
