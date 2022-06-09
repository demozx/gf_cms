package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
	"runtime"
	"time"
)

type cWelcome struct{}

var (
	Welcome = cWelcome{}
)

func (c *cWelcome) Index(ctx context.Context, req *backendApi.GetRuntimeInfoApiReq) (res *backendApi.GetRuntimeInfoApiRes, err error) {
	var cpu = service.Runtime().GetCpuInfo()
	var load = service.Runtime().GetLoadInfo()
	var mem = service.Runtime().GetMemInfo()
	var desk = service.Runtime().GetDiskInfo()
	var net = service.Runtime().GetNetInfo()
	serverStartAt := service.Runtime().GetServerStartAt()
	serverStartDuration := service.Util().FriendyTimeFormat(serverStartAt.Time(), time.Now())
	//g.Log().Info(ctx, "cpu", cpu)
	//g.Log().Info(ctx, "load", load)
	//g.Log().Info(ctx, "mem", mem)
	//g.Log().Info(ctx, "disk", desk)
	cPUNum := runtime.NumCPU()
	res = &backendApi.GetRuntimeInfoApiRes{
		Load:                load,
		Cpu:                 cpu,
		CPUNum:              cPUNum,
		Mem:                 mem,
		Disk:                desk,
		Net:                 net,
		ServerStartDuration: serverStartDuration,
	}
	return
}
