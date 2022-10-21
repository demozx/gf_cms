package backend

import "github.com/gogf/gf/v2/frame/g"

type ArticleMoveReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台登录"`
	Ids    string `json:"ids" in:"query" d:""  v:"required#必填项不能为空"  dc:"ids，英文逗号拼接"`
}
type ArticleMoveRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
