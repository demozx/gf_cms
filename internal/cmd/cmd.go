package cmd

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/logic/runtime"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/logic/viewBindFun"
	"gf_cms/internal/router"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
)

var (
	Main = gcmd.Command{
		Name:  util.Util().ProjectName(),
		Usage: util.Util().ProjectName(),
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 设置模板标签分隔符
			g.View().SetDelimiters("${", "}$")
			//设置服务启动时间
			runtime.Runtime().SetServerStartAt()
			s := g.Server(util.Util().ProjectName())
			//session使用redis
			_ = s.SetConfigWithMap(g.Map{
				// session一天过期
				"SessionMaxAge":  time.Hour * 24,
				"SessionStorage": gsession.NewStorageRedis(g.Redis(), util.Util().ProjectName()+":"+consts.AdminSessionKeyPrefix+":"),
			})
			//给模板视图全局绑定方法
			viewBindFun.ViewBindFun().Register()
			//路由
			router.Register(s)
			s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
				404: func(r *ghttp.Request) {
					// 输出pc端404页面
					_ = r.Response.WriteTpl(consts.Pc404Page)
				},
				403: func(r *ghttp.Request) {
					// 只清除缓冲区数据
					r.Response.ClearBuffer()
				},
				500: func(r *ghttp.Request) {
					err = r.Response.WriteTpl("tpl/error.html", g.Map{
						"code":    500,
						"message": "出错了，请稍后重试",
					})
				},
			})
			s.Run()
			return nil
		},
	}
)
