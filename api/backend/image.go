package backend

import "github.com/gogf/gf/v2/frame/g"

type ImageMoveReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"图集移动"`
	StrIds string `json:"str_ids" in:"query" d:""  v:"required#必填项不能为空"  dc:"str_ids，英文逗号拼接"`
}
type ImageMoveRes struct{}

type ImageAddReq struct {
	ChannelId int `json:"channel_id" in:"query" d:"0"  v:""  dc:"频道ID"`
	g.Meta    `tags:"Backend" method:"get" summary:"图集新增"`
}
type ImageAddRes struct{}

type ImageEditReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"图集编辑"`
	Id     int `json:"id" in:"query" d:"0"  v:"required#图集ID必填"  dc:"文章ID"`
}
type ImageEditRes struct{}
