package backend

import "github.com/gogf/gf/v2/frame/g"

type WelcomeReq struct {
	g.Meta `tags:"Backend" method:"get" summary:"后台欢迎页面"`
}
type WelcomeRes struct{}
