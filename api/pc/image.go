package pc

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ImageListReq struct {
	g.Meta    `tags:"Pc" method:"get" summary:"pc图集列表"`
	ChannelId int `json:"id" dc:"图集栏目id" d:"1"`
	model.PageSizeReq
}
type ImageListRes struct {
	List  []*model.ImageListItem
	Page  int `json:"page" description:"分页码"`
	Size  int `json:"size" description:"分页数量"`
	Total int `json:"total" description:"数据总数"`
}

type ImageDetailReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"pc图集详情"`
	Id     int `json:"id" dc:"图集id"`
}
type ImageDetailRes struct{}
