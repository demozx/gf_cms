package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleMoveReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台登录"`
	StrIds string `json:"str_ids" in:"query" d:""  v:"required#必填项不能为空"  dc:"str_ids，英文逗号拼接"`
}
type ArticleMoveRes struct{}
