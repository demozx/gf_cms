package mobile

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `tags:"Mobile" method:"get" summary:"Mobile首页"`
}
type IndexRes struct{}
