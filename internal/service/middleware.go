package service

import (
	"gf_cms/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type middlewareService struct{}

var middleware = middlewareService{}

func Middleware() *middlewareService {
	return &middleware
}

func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
	r.Response.Header().Set("Content-Type", "application/json;charset=utf-8")
}

func (s *middlewareService) Auth(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func (s *middlewareService) AdminAuthSession(r *ghttp.Request) {
	var adminSession, _ = r.Session.Get(consts.AdminSessionKeyPrefix)
	AdminPrefix, _ := g.Cfg().Get(r.Context(), "server.adminPrefix", "admin")
	if adminSession.IsEmpty() {
		AdminRoute := "/" + AdminPrefix.String()
		// 如果没有session且是get请求且当前页面不是后台入口，跳转到后台入口
		r.Response.RedirectTo(AdminRoute + "/admin/login")
	}
	r.Middleware.Next()
}
