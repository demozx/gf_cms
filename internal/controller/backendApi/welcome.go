package backendApi

import (
	"context"
	"fmt"
	"gf_cms/api/backendApi"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/gvar"
)

type cWelcome struct{}

var (
	Welcome = cWelcome{}
)

func (c *cWelcome) Index(ctx context.Context, req *backendApi.GetRuntimeInfoApiReq) (res *backendApi.GetRuntimeInfoApiRes, err error) {
	var cpu = service.Runtime().GetCpuInfo()[0]
	var load = service.Runtime().GetLoadInfo()
	var mem = service.Runtime().GetMemInfo()
	var desk = service.Runtime().GetDiskInfo()
	fmt.Println("desk", desk)

	res = &backendApi.GetRuntimeInfoApiRes{
		Load: gvar.New(load.Load1).String(),
		Cpu:  gvar.New(cpu).String(),
		Mem:  gvar.New(mem.UsedPercent).String(),
		Disk: gvar.New(desk.UsedPercent).String(),
	}
	return
}
