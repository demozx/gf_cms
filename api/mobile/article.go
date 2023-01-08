package mobile

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta    `tags:"Mobile" method:"get" summary:"mobile文章列表"`
	ChannelId int `json:"id" dc:"文章栏目id" d:"1"`
	Page      int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size      int `json:"size" in:"query" d:"15" v:"max:100#分页数量最大100条" dc:"分页数量，最大100"`
}
type ArticleListRes struct {
	List  []*model.ArticleListItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type ArticleDetailReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"mobile文章详情"`
	Id     int `json:"id" dc:"文章id"`
}
type ArticleDetailRes struct{}
