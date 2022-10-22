package backend

import "github.com/gogf/gf/v2/frame/g"

type SettingReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台设置"`
}
type SettingRes struct{}
