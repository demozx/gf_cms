package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"runtime"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Index(ctx context.Context, req *backend.WelcomeReq) (res *backend.WelcomeRes, err error) {
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("welcome/index.html", g.Map{
		"project_name": service.ProjectName,
		"system_root":  service.SystemRoot,
		"host_info":    service.Runtime().GetHostInfo(),
		"cpu_info":     service.Runtime().GetCpuInfo(),
		"go_version":   runtime.Version(),
		"go_root":      runtime.GOROOT(),
		"CPU_num":      runtime.NumCPU(),
	})
	return
}
