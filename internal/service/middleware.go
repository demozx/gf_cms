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

func (s *middlewareService) BackendAuthSession(r *ghttp.Request) {
	var adminSession, _ = r.Session.Get(consts.AdminSessionKeyPrefix)
	backendPrefix := Util().BackendPrefix()
	if adminSession.IsEmpty() {
		adminRoute := "/" + backendPrefix
		// 如果没有session且是get请求且当前页面不是后台入口，跳转到后台入口
		r.Response.RedirectTo(adminRoute + "/admin/login")
	}
	r.Middleware.Next()
}

// BackendCheckPolicy 检测用户有无某个请求权限
func (s *middlewareService) BackendCheckPolicy(r *ghttp.Request) {
	accountId := Middleware().GetAdminUserID(r)
	obj := CasbinPolicy().ObjBackend()
	act := r.Router.Uri
	g.Log().Info(Ctx, "act", act)
	var backendPrefix = Util().BackendPrefix()
	var backendAllMenus = Menu().BackendAll()
	var backendMyMenus = Menu().BackendMy(accountId)
	var routeHit = false
	for _, menuGroup := range backendAllMenus {
		for _, children := range menuGroup.Children {
			if act == "/"+backendPrefix+children.Route {
				routeHit = true
			}
		}
	}
	g.Log().Info(Ctx, "routeHit："+act, routeHit)
	if routeHit == false {
		//路由不在权限中，不拦截
		r.Middleware.Next()
	} else {
		var routePermission = ""
		for _, menu := range backendMyMenus {
			for _, children := range menu.Children {
				if "/"+backendPrefix+children.Route == act {
					routePermission = children.Permission
					g.Log().Info(Ctx, "路由"+act+"的权限是："+children.Permission)
				}
			}
		}
		if !CasbinPolicy().CheckByAccountId(accountId, obj, routePermission) {
			g.Log().Info(Ctx, "没有权限"+act)
			r.Response.WriteExit("没有权限")
		}
		r.Middleware.Next()
	}
}

// GetAdminUserID 获取后台用户ID
func (s *middlewareService) GetAdminUserID(r *ghttp.Request) string {
	var adminSession, _ = r.Session.Get(consts.AdminSessionKeyPrefix)
	var admin *entity.CmsAdmin
	err := adminSession.Scan(&admin)
	if err != nil {
		panic(err)
	}
	var res = g.Map{
		"id": admin.Id,
	}
	accountId := gvar.New(res["id"]).String()
	return accountId
}
