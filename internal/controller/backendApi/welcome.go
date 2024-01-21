package backendApi

import (
	"context"
	"gf_cms/api/backendApi"
	runtime2 "gf_cms/internal/logic/runtime"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gproc"
	"runtime"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

type cWelcome struct{}

var (
	Welcome = cWelcome{}
)

func (c *cWelcome) Index(ctx context.Context, req *backendApi.GetRuntimeInfoApiReq) (res *backendApi.GetRuntimeInfoApiRes, err error) {
	var cpu = runtime2.Runtime().GetCpuInfo()
	var load = runtime2.Runtime().GetLoadInfo()
	var mem = runtime2.Runtime().GetMemInfo()
	var desk = runtime2.Runtime().GetDiskInfo()
	var net = runtime2.Runtime().GetNetInfo()
	serverStartAt := runtime2.Runtime().GetServerStartAt()
	serverStartDuration := util.Util().FriendyTimeFormat(serverStartAt.Time(), time.Now())
	cpuNum := runtime.NumCPU()
	var loadPercent = gconv.Float32(load.Load1) * 100 / gconv.Float32(cpuNum) * gconv.Float32(0.5)
	if loadPercent > 100 {
		loadPercent = gconv.Float32(100)
	}
	var numGoroutine = runtime.NumGoroutine()
	var mysqlProcessNum = runtime2.Runtime().MysqlProcessNum()
	var mySqlMaxConnectionsNum = runtime2.Runtime().MySqlMaxConnectionsNum()
	var mySqlCurrConnectionsNum = runtime2.Runtime().MySqlCurrConnectionsNum()
	var redisMaxClientsNum = runtime2.Runtime().RedisMaxClientsNum()
	var redisConnectedClientsNum = runtime2.Runtime().RedisConnectedClientsNum()

	res = &backendApi.GetRuntimeInfoApiRes{
		Load:                     load,
		LoadPercent:              loadPercent,
		Cpu:                      cpu,
		CPUNum:                   cpuNum,
		Mem:                      mem,
		Disk:                     desk,
		Net:                      net,
		ServerStartDuration:      serverStartDuration,
		NumGoroutine:             numGoroutine,
		MysqlProcessNum:          mysqlProcessNum,
		MySqlMaxConnectionsNum:   mySqlMaxConnectionsNum,
		MySqlCurrConnectionsNum:  mySqlCurrConnectionsNum,
		RedisMaxClientsNum:       redisMaxClientsNum,
		RedisConnectedClientsNum: redisConnectedClientsNum,
		Pid:                      gproc.Pid(),
	}
	return
}

func (c *cWelcome) RestartServer(ctx context.Context, req *backendApi.RestartServerReq) (res *backendApi.RestartServerRes, err error) {
	if service.Util().GetConfig("server.graceful") != "true" {
		return nil, gerror.New("未开启平滑重启特性")
	}
	err = ghttp.RestartAllServer(ctx, "")
	if err != nil {
		return nil, err
	}
	service.Response().SuccessJsonDefaultMessage(ctx, "重启成功")
	return
}
