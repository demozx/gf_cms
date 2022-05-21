package service

import (
	"gf_cms/internal/consts"
	"gf_cms/internal/model/entity"
	"github.com/gogf/gf/v2/container/gvar"
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
	adminPrefix := Util().AdminPrefix()
	if adminSession.IsEmpty() {
		adminRoute := "/" + adminPrefix
		// 如果没有session且是get请求且当前页面不是后台入口，跳转到后台入口
		r.Response.RedirectTo(adminRoute + "/admin/login")
	}
	r.Middleware.Next()
}

func (s *middlewareService) AdminCheckPolicy(r *ghttp.Request) {
	var adminSession, _ = r.Session.Get(consts.AdminSessionKeyPrefix)
	var admin *entity.CmsAdmin
	err := adminSession.Scan(&admin)
	if err != nil {
		panic(err)
	}
	var res = g.Map{
		"id": admin.Id,
	}
	if !Policy().CheckByAccountId(gvar.New(res["id"]).String(), Policy().ObjBackend(), r.Router.Uri) {
		r.Response.WriteExit("没有权限")
	}
	r.Middleware.Next()
}
