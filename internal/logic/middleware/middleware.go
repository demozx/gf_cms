package middleware

import (
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/auth"
	"gf_cms/internal/logic/casbinPolicy"
	"gf_cms/internal/logic/menu"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/model/entity"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"

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

// BackendCheckPolicy 检测后台页面用户有无某个请求权限
func (s *sMiddleware) BackendCheckPolicy(r *ghttp.Request) {
	accountId := Middleware().GetBackendUserID(r)
	obj := casbinPolicy.CasbinPolicy().ObjBackend()
	act := r.Router.Uri
	//g.Log().Notice(util.Ctx, "act", act)
	var backendPrefix = util.Util().BackendPrefix()
	var backendViewMenus = menu.Menu().BackendView()
	var backendMyMenus = menu.Menu().BackendMyMenu(accountId)
	var routeHit = false
	for _, menuGroup := range backendViewMenus {
		for _, children := range menuGroup.Children {
			if act == "/"+backendPrefix+children.Route {
				routeHit = true
			}
		}
	}
	//g.Log().Notice(util.Ctx, "backendRouteHit："+act, routeHit)
	if routeHit == false {
		//路由不在权限中，不拦截
		r.Middleware.Next()
	} else {
		var routePermission = ""
		for _, menu := range backendMyMenus {
			for _, children := range menu.Children {
				if "/"+backendPrefix+children.Route == act {
					routePermission = children.Permission
					g.Log().Notice(util.Ctx, "路由"+act+"的权限是："+children.Permission)
				}
			}
		}
		if !casbinPolicy.CasbinPolicy().CheckByAccountId(accountId, obj, routePermission) {
			g.Log().Warning(util.Ctx, "没有权限"+act)
			err := r.Response.WriteTpl("tpl/error.html", g.Map{
				"code":    401,
				"message": "无权访问",
			})
			if err != nil {
				return
			}
		} else {
			r.Middleware.Next()
		}
	}
}

// BackendApiCheckPolicy 检测后台接口用户有无某个请求权限
func (s *sMiddleware) BackendApiCheckPolicy(r *ghttp.Request) {
	//先取session，没有session走jwt
	var sessionInfo, _ = r.Session.Get(consts.AdminSessionKeyPrefix)
	var adminSession *model.AdminSession
	err := sessionInfo.Scan(&adminSession)
	if err != nil {
		return
	}
	var accountId string
	if adminSession != nil {
		accountId = gconv.String(adminSession.Id)
	} else {
		Middleware().Auth(r)
		accountId = gconv.String(auth.Auth().JWTAuth().GetIdentity(r.GetCtx()))
	}
	obj := casbinPolicy.CasbinPolicy().ObjBackendApi()
	act := r.Router.Uri
	g.Log().Notice(util.Ctx, "act", act)
	var backendPrefix = util.Util().BackendApiPrefix()
	var backendApiMenus = menu.Menu().BackendApi()
	var backendMyApis = menu.Menu().BackendMyApi(accountId)
	//g.Dump("backendApiMenus", backendApiMenus)
	//g.Dump("backendMyApis", backendMyApis)
	var routeHit = false
	var ruleName = ""
	for _, menuGroup := range backendApiMenus {
		for _, children := range menuGroup.Children {
			if act == "/"+backendPrefix+children.Route {
				routeHit = true
				ruleName = children.Title
			}
		}
	}
	g.Log().Notice(util.Ctx, "backendApiRouteHit："+act, routeHit)
	if routeHit == false {
		//路由不在权限中，不拦截
		r.Middleware.Next()
	} else {
		var routePermission = ""
		for _, menu := range backendMyApis {
			for _, children := range menu.Children {
				if "/"+backendPrefix+children.Route == act {
					routePermission = children.Permission
					g.Log().Notice(util.Ctx, "路由"+act+"的权限是："+children.Permission)
				}
			}
		}
		if !casbinPolicy.CasbinPolicy().CheckByAccountId(accountId, obj, routePermission) {
			g.Log().Warning(util.Ctx, "没有权限"+act)
			r.Response.WriteJsonExit(g.Map{
				"code":    401,
				"message": "没有权限操作 " + ruleName,
			})
		}
		r.Middleware.Next()
	}
}

// GetBackendUserID 获取后台用户ID
func (s *sMiddleware) GetBackendUserID(r *ghttp.Request) string {
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
