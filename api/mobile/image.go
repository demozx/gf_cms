package mobile

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ImageListReq struct {
	g.Meta    `tags:"Mobile" method:"get" summary:"mobile图集列表"`
	ChannelId int `json:"id" dc:"图集栏目id" d:"1"`
	Page      int `json:"page" in:"query" d:"1"  v:"min:0#分页号码错误"     dc:"分页号码，默认1"`
	Size      int `json:"size" in:"query" d:"10" v:"max:100#分页数量最大100条" dc:"分页数量，最大100"`
}
type ImageListRes struct {
	List  []*model.ImageListItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type ImageDetailReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"mobile图集详情"`
	Id     int `json:"id" dc:"图集id"`
}
type ImageDetailRes struct{}
