package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/util/gconv"
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
	cpuNum := runtime.NumCPU()
	var loadPercent = gconv.Float32(load.Load1) * 100 / gconv.Float32(cpuNum) * gconv.Float32(0.5)
	if loadPercent > 100 {
		loadPercent = gconv.Float32(100)
	}
	res = &backendApi.GetRuntimeInfoApiRes{
		Load:                load,
		LoadPercent:         loadPercent,
		Cpu:                 cpu,
		CPUNum:              cpuNum,
		Mem:                 mem,
		Disk:                desk,
		Net:                 net,
		ServerStartDuration: serverStartDuration,
	}
	return
}
