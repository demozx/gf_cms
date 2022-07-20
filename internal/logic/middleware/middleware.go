package middleware

import (
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/auth"
	"gf_cms/internal/logic/casbinPolicy"
	"gf_cms/internal/logic/menu"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

var middleware = sMiddleware{}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}

func Middleware() *sMiddleware {
	return &middleware
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
	r.Response.Header().Set("Content-Type", "application/json;charset=utf-8")
}

func (s *sMiddleware) Auth(r *ghttp.Request) {
	auth.Auth().JWTAuth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

func (s *sMiddleware) BackendAuthSession(r *ghttp.Request) {
	var adminSession, _ = r.Session.Get(consts.AdminSessionKeyPrefix)
	backendPrefix := util.Util().BackendPrefix()
	if adminSession.IsEmpty() {
		adminRoute := "/" + backendPrefix
		// 如果没有session且是get请求且当前页面不是后台入口，跳转到后台入口
		r.Response.RedirectTo(adminRoute + "/admin/login")
	}
	r.Middleware.Next()
}

// BackendCheckPolicy 检测用户有无某个请求权限
func (s *sMiddleware) BackendCheckPolicy(r *ghttp.Request) {
	accountId := Middleware().GetAdminUserID(r)
	obj := casbinPolicy.CasbinPolicy().ObjBackend()
	act := r.Router.Uri
	g.Log().Info(util.Ctx, "act", act)
	var backendPrefix = util.Util().BackendPrefix()
	var backendAllMenus = menu.Menu().BackendAll()
	var backendMyMenus = menu.Menu().BackendMy(accountId)
	var routeHit = false
	for _, menuGroup := range backendAllMenus {
		for _, children := range menuGroup.Children {
			if act == "/"+backendPrefix+children.Route {
				routeHit = true
			}
		}
	}
	g.Log().Info(util.Ctx, "routeHit："+act, routeHit)
	if routeHit == false {
		//路由不在权限中，不拦截
		r.Middleware.Next()
	} else {
		var routePermission = ""
		for _, menu := range backendMyMenus {
			for _, children := range menu.Children {
				if "/"+backendPrefix+children.Route == act {
					routePermission = children.Permission
					g.Log().Info(util.Ctx, "路由"+act+"的权限是："+children.Permission)
				}
			}
		}
		if !casbinPolicy.CasbinPolicy().CheckByAccountId(accountId, obj, routePermission) {
			g.Log().Info(util.Ctx, "没有权限"+act)
			r.Response.WriteExit("没有权限")
		}
		r.Middleware.Next()
	}
}

// GetAdminUserID 获取后台用户ID
func (s *sMiddleware) GetAdminUserID(r *ghttp.Request) string {
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
