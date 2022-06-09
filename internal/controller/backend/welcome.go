package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"runtime"
	"time"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Index(ctx context.Context, req *backend.WelcomeReq) (res *backend.WelcomeRes, err error) {
	ip, _ := service.Util().GetLocalIP()
	serverAddress, _ := g.Config().Get(ctx, "server.address")
	serverStartAt := service.Runtime().GetServerStartAt()
	serverStartDuration := service.Util().FriendyTimeFormat(serverStartAt.Time(), time.Now())
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("backend/welcome/index.html", g.Map{
		"project_name":          service.ProjectName,
		"system_root":           service.SystemRoot,
		"host_info":             service.Runtime().GetHostInfo(),
		"cpu_info":              service.Runtime().GetCpuInfo(),
		"go_version":            runtime.Version(),
		"go_root":               runtime.GOROOT(),
		"CPU_num":               runtime.NumCPU(),
		"ip":                    ip,
		"server_address":        serverAddress,
		"server_start_at":       serverStartAt,
		"server_start_duration": serverStartDuration,
	})
	return
}
