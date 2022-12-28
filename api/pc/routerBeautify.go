package pc

import "github.com/gogf/gf/v2/frame/g"

type RouterBeautifyAboutReq struct {
	g.Meta `tags:"Pc" method:"get" summary:"pc路由美化-关于我们"`
	Id     int `json:"id" d:"8"`
}

type RouterBeautifyRes struct{}
