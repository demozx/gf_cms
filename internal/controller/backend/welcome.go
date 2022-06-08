package backend

import (
	"context"
	"gf_cms/api/backend"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"runtime"
)

var (
	Welcome = cWelcome{}
)

type cWelcome struct{}

func (c *cWelcome) Index(ctx context.Context, req *backend.WelcomeReq) (res *backend.WelcomeRes, err error) {
	diskInfo := service.Runtime().GetDiskInfo()
	diskInfo.All = gvar.New(gvar.New(diskInfo.All).Int() / 1024).String()
	diskInfo.Used = gvar.New(gvar.New(diskInfo.Used).Int() / 1024).String()
	memInfo := service.Runtime().GetMemInfo()
	memInfo.Total = gvar.New(gvar.New(memInfo.Total).Int() / 1024 / 1024 / 1024).String()
	memInfo.Used = gvar.New(gvar.New(memInfo.Used).Int() / 1024 / 1024 / 1024).String()
	_ = g.RequestFromCtx(ctx).Response.WriteTpl("welcome/index.html", g.Map{
		"project_name": service.ProjectName,
		"system_root":  service.SystemRoot,
		"host_info":    service.Runtime().GetHostInfo(),
		"cpu_info":     service.Runtime().GetCpuInfo(),
		"disk_info":    diskInfo,
		"mem_info":     memInfo,
		"go_version":   runtime.Version(),
		"go_root":      runtime.GOROOT(),
	})
	return
}
