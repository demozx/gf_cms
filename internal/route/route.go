package route

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Register(s *ghttp.Server) {
	//后台路由
	backendViewHandle(s)
	backendApiHandle(s)
	//前台路由
	webViewHandle(s)
}
