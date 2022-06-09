package cmd

import (
	"context"
	"gf_cms/internal/consts"
	"gf_cms/internal/route"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			//设置服务启动时间
			service.Runtime().SetServerStartAt()

			s := g.Server()

			//session使用redis
			_ = s.SetConfigWithMap(g.Map{
				// session一天过期
				"SessionMaxAge":  time.Hour * 24,
				"SessionStorage": gsession.NewStorageRedis(g.Redis(), service.Util().ProjectName()+":"+consts.AdminSessionKeyPrefix+":"),
			})

			//给模板视图全局绑定方法
			service.ViewBindFun().Register()

			//路由
			route.Register(s)

			s.Run()

			return nil
		},
	}
)
