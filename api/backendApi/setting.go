package backendApi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SettingSaveApiReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"保存后台设置"`
}
type SettingSaveApiRes struct {
}
