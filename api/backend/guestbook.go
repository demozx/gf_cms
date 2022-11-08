package backend

import (
	"gf_cms/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type GuestbookIndexReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"留言列表"`
	model.PageSizeReq
}
type GuestbookIndexRes struct{}
