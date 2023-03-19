package runtime

import (
	"fmt"
	"gf_cms/internal/logic/util"
	"gf_cms/internal/model"
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"gitlab.com/tingshuo/go-diskstate/diskstate"
)

type sRuntime struct{}

var (
	insRuntime = sRuntime{}
)

// 服务启动时间缓存key
var serverStartAtCacheKey = "serverStartAt"

func init() {
	service.RegisterRuntime(New())
}

func New() *sRuntime {
	return &sRuntime{}
}

func Runtime() *sRuntime {
	return &insRuntime
}

// GetCpuInfo CPU信息
func (*sRuntime) GetCpuInfo() model.Cpu {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Println("get cpu info failed: ", err)
	}
	var cpuInfo model.Cpu
	for _, ci := range cpuInfos {
		//fmt.Println(ci)
		cpuInfo.Cores = gvar.New(ci.Cores).Int()
		cpuInfo.ModelName = gvar.New(ci.ModelName).String()
		cpuInfo.Mhz = gvar.New(ci.Mhz).Int()
		cpuInfo.CurrMhz = gvar.New(ci.Mhz).Int()
	}
	percent, _ := cpu.Percent(0, false) // false表示CPU总使用率，true为单核
	if len(percent) > 0 {
		cpuInfo.UsedPercent = gvar.New(percent[0]).String()
	} else {
		cpuInfo.UsedPercent = ""
	}
	if runtime.GOOS == "linux" {
		command := exec.Command("bash", "-c", `cat /proc/cpuinfo |grep MHz`)
		output, _ := command.CombinedOutput()
		split := gstr.Split(gconv.String(output), "\n")
		if len(split) > 0 {
			from := garray.NewStrArrayFrom(split).Sort()
			value, found := from.Get(from.Len() - 1)
			if found {
				row := gstr.Split(value, "\t\t:")
				if len(row) > 1 {
					cpuInfo.CurrMhz = gconv.Int(gstr.Trim(row[1]))
				}
			}
		}
	}

	return cpuInfo
}

// GetLoadInfo 负载信息
func (*sRuntime) GetLoadInfo() model.Load {
	info, err := load.Avg()
	if err != nil {
		fmt.Println("load.Avg() failed: ", err)
	}
	var loadInfo model.Load

	loadInfo.Load1 = gvar.New(info.Load1).String()
	loadInfo.Load5 = gvar.New(info.Load5).String()
	loadInfo.Load15 = gvar.New(info.Load15).String()
	return loadInfo
}

// GetMemInfo 内存信息
func (*sRuntime) GetMemInfo() model.Mem {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("mem.VirtualMemory() failed: ", err)
	}
	var memIn model.Mem
	memIn.Total = gvar.New(memInfo.Total).String()
	memIn.Available = gvar.New(memInfo.Available).String()
	memIn.Used = gvar.New(memInfo.Used).String()
	memIn.UsedPercent = gvar.New(memInfo.UsedPercent).String()
	memIn.Free = gvar.New(memInfo.Free).String()
	memIn.Active = gvar.New(memInfo.Active).String()
	memIn.Inactive = gvar.New(memInfo.Inactive).String()
	memIn.Wired = gvar.New(memInfo.Wired).String()
	return memIn
}

// GetHostInfo 主机信息
func (*sRuntime) GetHostInfo() model.Host {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("host.Info() failed: ", err)
	}
	var hostIn model.Host
	hostIn.Hostname = hostInfo.Hostname
	hostIn.Uptime = gvar.New(hostInfo.Uptime).String()
	hostIn.BootTime = gvar.New(hostInfo.BootTime).String()
	hostIn.Procs = gvar.New(hostInfo.Procs).String()
	hostIn.OS = hostInfo.OS
	hostIn.Platform = hostInfo.Platform
	hostIn.PlatformFamily = hostInfo.PlatformFamily
	hostIn.PlatformVersion = hostInfo.PlatformVersion
	hostIn.KernelVersion = hostInfo.KernelVersion
	hostIn.KernelArch = hostInfo.KernelArch
	hostIn.VirtualizationSystem = hostInfo.VirtualizationSystem
	hostIn.VirtualizationRole = hostInfo.VirtualizationRole
	hostIn.HostID = hostInfo.HostID

	return hostIn
}

// GetDiskInfo 磁盘信息
func (*sRuntime) GetDiskInfo() model.Disk {
	//parts, err := disk.Partitions(true)
	//if err != nil {
	//	fmt.Println("get disk partitions failed: ", err)
	//}
	//for _, part := range parts {
	//	partInfo, err := disk.Usage(part.Mountpoint)
	//	if err != nil {
	//		fmt.Println("get part stat failed: ", err)
	//		return nil
	//	}
	//	return partInfo
	//}
	//return nil
	state := diskstate.DiskUsage("/")
	var disk model.Disk
	disk.All = gvar.New(state.All / diskstate.MB).String()
	disk.Free = gvar.New(state.Free / diskstate.MB).String()
	disk.Available = gvar.New(state.Available / diskstate.MB).String()
	disk.Used = gvar.New(state.Used / diskstate.MB).String()
	disk.Usage = gvar.New(100 * state.Used / state.All).String()
	return disk
}

// GetNetInfo 网络信息
func (*sRuntime) GetNetInfo() model.Net {
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("get net io counters failed: ", err)
	}
	var (
		kbsSent = 0
		kbsRecv = 0
	)
	//fmt.Printf("%v", netIOs)
	for _, netIO := range netIOs {
		if strings.HasPrefix(netIO.Name, "en") {
			//fmt.Println(netIO) // 打印每张网卡信息
			kbsSent = int(netIO.BytesSent/1024) + kbsSent
			kbsRecv = int(netIO.BytesRecv/1024) + kbsRecv
		}
	}
	//fmt.Println(kbsSent, kbsRecv)

	var kbsSentCacheKey = util.Util().ProjectName() + ":net_info:kbs_sent"
	var kbsRecvCacheKey = util.Util().ProjectName() + ":net_info:kbs_recv"
	var kbsTimeCacheKey = util.Util().ProjectName() + ":net_info:kbs_time"
	conn, err := g.Redis().Conn(util.Ctx)
	kbsSentCached, _ := conn.Do(util.Ctx, "GET", kbsSentCacheKey)
	kbsRecvCached, _ := conn.Do(util.Ctx, "GET", kbsRecvCacheKey)
	kbsTimeCached, _ := conn.Do(util.Ctx, "GET", kbsTimeCacheKey)

	//fmt.Println("kbsSentCached", kbsSentCached)
	//fmt.Println("kbsRecvCached", kbsRecvCached)
	//fmt.Println("kbsTimeCached", kbsTimeCached)

	conn.Do(util.Ctx, "SET", kbsSentCacheKey, kbsSent)
	conn.Do(util.Ctx, "SET", kbsRecvCacheKey, kbsRecv)
	conn.Do(util.Ctx, "SET", kbsTimeCacheKey, time.Now().Unix())
	var netInfo model.Net
	var seconds = gvar.New(time.Now().Unix()).Int() - kbsTimeCached.Int()
	if seconds > 0 {
		netInfo.KbsSent = (kbsSent - kbsSentCached.Int()) / seconds
		netInfo.KbsRecv = (kbsRecv - kbsRecvCached.Int()) / seconds
	} else {
		netInfo.KbsSent = 0
		netInfo.KbsRecv = 0
	}
	defer conn.Close(util.Ctx)
	return netInfo
}

// SetServerStartAt 设置服务启动时间
func (*sRuntime) SetServerStartAt() bool {
	gcache.Set(util.Ctx, serverStartAtCacheKey, time.Now().Unix(), 0)
	return true
}

// GetServerStartAt 获取服务启动时间
func (*sRuntime) GetServerStartAt() *gvar.Var {
	get, _ := gcache.Get(util.Ctx, serverStartAtCacheKey)
	return get
}

// MysqlProcessNum MySql进程数
func (*sRuntime) MysqlProcessNum() int {
	query, err := g.DB().Query(util.Ctx, "show full processlist")
	if err != nil {
		return 0
	}
	return query.Len()
}

// MySqlMaxConnectionsNum MySql最大连接数
func (*sRuntime) MySqlMaxConnectionsNum() int {
	query, err := g.DB().Query(util.Ctx, "show variables like '%max_connections%'")
	if err != nil {
		return 0
	}
	return gconv.Int(query.Array("Value")[0])
}

// MySqlCurrConnectionsNum MySql当前连接数
func (*sRuntime) MySqlCurrConnectionsNum() int {
	query, err := g.DB().Query(util.Ctx, "show global status like 'Max_used_connections'")
	if err != nil {
		return 0
	}
	return gconv.Int(query.Array("Value")[0])
}

// RedisMaxClientsNum Redis最大连接数
func (*sRuntime) RedisMaxClientsNum() int {
	do, _ := g.Redis().Do(util.Ctx, "config", "get", "maxclients")
	return gconv.Int(do.Array()[1])
}

// RedisConnectedClientsNum 获取Redis当前连接数
func (*sRuntime) RedisConnectedClientsNum() int {
	do, _ := g.Redis().Do(util.Ctx, "info", "clients")
	res := do.String()
	res = gstr.Nl2Br(res)
	split := gstr.Split(res, "<br>")
	fullConnectedClients := split[1]
	connectedClients := gstr.Replace(fullConnectedClients, "connected_clients:", "")
	return gconv.Int(connectedClients)
}
