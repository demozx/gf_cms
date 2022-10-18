package router

import (
	"gf_cms/internal/controller/backendApi"
	"gf_cms/internal/logic/middleware"
	"gf_cms/internal/logic/util"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

//后台api路由分组
func backendApiHandle(s *ghttp.Server) {
	var backendApiGroup = util.Util().BackendApiGroup()
	s.Group(backendApiGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			middleware.Middleware().CORS,
		)
		group.ALLMap(g.Map{
			"/captcha/get": backendApi.Captcha.Get,
			"/admin/login": backendApi.Admin.Login,
		})
	})
	s.Group(backendApiGroup, func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareHandlerResponse,
			middleware.Middleware().BackendApiCheckPolicy,
		)
		group.ALLMap(g.Map{
			"/welcome/index": backendApi.Welcome.Index, //欢迎页面
			"/setting/save":  backendApi.Setting.Save,  //保存系统设置
			/*管理员*/
			"/admin/logout":       backendApi.Admin.Logout,      //退出登录
			"/admin/clear_cache":  backendApi.Admin.ClearCache,  //清理缓存
			"/admin/add":          backendApi.Admin.Add,         //添加管理员
			"/admin/edit":         backendApi.Admin.Edit,        //编辑管理员
			"/admin/status":       backendApi.Admin.Status,      //启动禁用
			"/admin/delete":       backendApi.Admin.Delete,      //删除
			"/admin/delete_batch": backendApi.Admin.DeleteBatch, //批量删除
			/*角色*/
			"/role/status":       backendApi.Role.Status,      //启用禁用
			"/role/delete":       backendApi.Role.Delete,      //删除
			"/role/delete_batch": backendApi.Role.DeleteBatch, //批量删除
			"/role/add":          backendApi.Role.Add,         //添加角色
			"/role/edit":         backendApi.Role.Edit,        //添加角色
			/*栏目分类*/
			"/channel/index":  backendApi.Channel.Index,  //栏目分类列表
			"/channel/status": backendApi.Channel.Status, //启用禁用
			"/channel/delete": backendApi.Channel.Delete, //删除
			"/channel/add":    backendApi.Channel.Add,    //添加
			"/channel/edit":   backendApi.Channel.Edit,   //编辑
			/*文章模型*/
			"/article/index": backendApi.Article.Index, // 文章列表
			"/article/sort":  backendApi.Article.Sort,  // 文章排序

			/*上传*/
			"/upload/single_file":  backendApi.Upload.SingleFile,  //文件上传
			"/upload/single_image": backendApi.Upload.SingleImage, //图片上传
			"/upload/single_video": backendApi.Upload.SingleVideo, //视频上传
		})
	})
}
