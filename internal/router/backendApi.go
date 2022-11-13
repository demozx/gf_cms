package router

import (
	"gf_cms/internal/controller/backendApi"
	"gf_cms/internal/logic/middleware"
	"gf_cms/internal/logic/util"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 后台api路由分组
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
			"/article/index":  backendApi.Article.Index,  // 文章列表
			"/article/sort":   backendApi.Article.Sort,   // 文章排序
			"/article/flag":   backendApi.Article.Flag,   // flag
			"/article/status": backendApi.Article.Status, // 审核状态
			"/article/delete": backendApi.Article.Delete, // 删除
			"/article/move":   backendApi.Article.Move,   // 移动
			"/article/add":    backendApi.Article.Add,    // 新增
			"/article/edit":   backendApi.Article.Edit,   // 编辑
			/*回收站*/
			"/recycle_bin/article_list":          backendApi.RecycleBin.ArticleList,         //文章列表
			"/recycle_bin/article_batch_destroy": backendApi.RecycleBin.ArticleBatchDestroy, //文章批量永久删除
			"/recycle_bin/article_batch_restore": backendApi.RecycleBin.ArticleBatchRestore, //文章批量恢复
			/*广告管理*/
			"/ad_channel/index":     backendApi.AdChannel.Index,    //广告分类列表
			"/ad_channel/add":       backendApi.AdChannel.Add,      //添加广告分类
			"/ad_channel/edit":      backendApi.AdChannel.Edit,     //编辑广告分类
			"/ad_channel/delete":    backendApi.AdChannel.Delete,   //删除广告分类
			"/ad_channel/sort":      backendApi.AdChannel.Sort,     //广告分类排序
			"/ad_list/index":        backendApi.AdList.Index,       //广告列表
			"/ad_list/add":          backendApi.AdList.Add,         //添加广告
			"/ad_list/edit":         backendApi.AdList.Edit,        //编辑广告
			"/ad_list/delete":       backendApi.AdList.Delete,      //删除广告
			"/ad_list/batch_status": backendApi.AdList.BatchStatus, //批量修改广告状态
			"/ad_list/sort":         backendApi.AdList.Sort,        //批量广告排序
			/*留言管理*/
			"/guestbook/status":       backendApi.Guestbook.Status,      //修改留言状态
			"/guestbook/batch_delete": backendApi.Guestbook.BatchDelete, //批量删除留言
			/*友情链接*/
			"/friendly_link/status":       backendApi.FriendlyLink.Status,      //修改状态
			"/friendly_link/add":          backendApi.FriendlyLink.Add,         //添加
			"/friendly_link/edit":         backendApi.FriendlyLink.Edit,        //编辑
			"/friendly_link/sort":         backendApi.FriendlyLink.Sort,        //排序
			"/friendly_link/batch_delete": backendApi.FriendlyLink.BatchDelete, //批量删除
			/*快捷方式*/
			"/shortcut/add":          backendApi.Shortcut.Add,         //添加
			"/shortcut/edit":         backendApi.Shortcut.Edit,        //编辑
			"/shortcut/batch_delete": backendApi.Shortcut.BatchDelete, //编辑
			"/shortcut/sort":         backendApi.Shortcut.Sort,        //排序

			/*上传*/
			"/upload/single_file":  backendApi.Upload.SingleFile,  //文件上传
			"/upload/single_image": backendApi.Upload.SingleImage, //图片上传
			"/upload/single_video": backendApi.Upload.SingleVideo, //视频上传
		})
	})
}
