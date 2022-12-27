package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(s *ghttp.Server) {
	//后台路由
	backendViewHandle(s)
	backendApiHandle(s)
	//前台pc路由
	pcViewHandle(s)
	pcApiHandle(s)
}
