package backend

import (
	"context"
	"gf_cms/api/backend"
	runtime2 "gf_cms/internal/logic/runtime"
	"gf_cms/internal/logic/util"
	"github.com/gogf/gf/v2"
	"runtime"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Index(ctx context.Context, req *backend.WelcomeReq) (res *backend.WelcomeRes, err error) {
	ip, _ := util.Util().GetLocalIP()
	serverAddress, _ := g.Config().Get(ctx, "server.address")
	serverStartAt := runtime2.Runtime().GetServerStartAt()
	serverStartDuration := util.Util().FriendyTimeFormat(serverStartAt.Time(), time.Now())
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("backend/welcome/index.html", g.Map{
		"project_name":          util.ProjectName,
		"system_root":           util.SystemRoot,
		"host_info":             runtime2.Runtime().GetHostInfo(),
		"cpu_info":              runtime2.Runtime().GetCpuInfo(),
		"go_version":            runtime.Version(),
		"gf_version":            gf.VERSION,
		"go_root":               runtime.GOROOT(),
		"cpu_num":               runtime.NumCPU(),
		"ip":                    ip,
		"server_address":        serverAddress,
		"server_start_at":       serverStartAt,
		"server_start_duration": serverStartDuration,
	})
	return
}
