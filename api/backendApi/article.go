package backendApi

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleListReq struct {
	g.Meta    `tags:"BackendApi" method:"post" summary:"文章列表"`
	ChannelId int    `p:"channel_id" name:"channel_id" v:"" des:"频道ID" brief:"频道ID" arg:"true"`
	StartAt   string `p:"start_at" name:"start_at" v:"date-format:Y-m-d H:i:s" des:"开始时间" brief:"频道ID" arg:"true"`
	EndAt     string `p:"end_at" name:"end_at" v:"date-format:Y-m-d H:i:s" des:"开始时间" brief:"频道ID" arg:"true"`
	Keyword   string `p:"keyword" name:"keyword" v:"" des:"关键词"`
	model.PageSizeReq
}

type ArticleListRes struct {
	List  []model.ArticleListItem `p:"list" name:"list" des:"列表"`
	Page  int                     `p:"page" name:"page" des:"分页码"`
	Size  int                     `p:"size" name:"size" des:"分页数量"`
	Total int                     `p:"total" name:"total" des:"数据总数"`
}
