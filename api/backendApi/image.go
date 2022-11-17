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
