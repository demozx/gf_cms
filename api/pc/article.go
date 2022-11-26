package pc

import "github.com/gogf/gf/v2/frame/g"

type ArticleListReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"pc文章列表"`
	Id     int `json:"id" dc:"文章列表id" d:"1"`
}
type ArticleListRes struct{}

type ArticleDetailReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"pc文章详情"`
	Id     int `json:"id" dc:"文章id"`
}
type ArticleDetailRes struct{}
