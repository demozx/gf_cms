package service

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"gitlab.com/tingshuo/go-diskstate/diskstate"
	"strings"
	"time"
)

type sRuntime struct{}

type Load struct {
	Load1  string
	Load5  string
	Load15 string
}

type Cpu struct {
	Cores       int
	ModelName   string
	Mhz         int
	UsedPercent string
}

type Mem struct {
	Total       string
	Available   string
	Used        string
	UsedPercent string
	Free        string
	Active      string
	Inactive    string
	Wired       string
}

type Disk struct {
	All       string
	Free      string
	Available string
	Used      string
	Usage     string
}

type Host struct {
	Hostname             string
	Uptime               string
	BootTime             string
	Procs                string
	OS                   string
	Platform             string
	PlatformFamily       string
	PlatformVersion      string
	KernelVersion        string
	KernelArch           string
	VirtualizationSystem string
	VirtualizationRole   string
	HostID               string
}

type Net struct {
	KbsSent int
	KbsRecv int
}

var (
	insRuntime = sRuntime{}
)

//服务启动时间缓存key
var serverStartAtCacheKey = "serverStartAt"

func Runtime() *sRuntime {
	return &insRuntime
}

// GetCpuInfo CPU信息
func (*sRuntime) GetCpuInfo() Cpu {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Println("get cpu info failed: ", err)
	}
	var cpuInfo Cpu
	for _, ci := range cpuInfos {
		//fmt.Println(ci)
		cpuInfo.Cores = gvar.New(ci.Cores).Int()
		cpuInfo.ModelName = gvar.New(ci.ModelName).String()
		cpuInfo.Mhz = gvar.New(ci.Mhz).Int()
	}
	percent, _ := cpu.Percent(0, false) // false表示CPU总使用率，true为单核
	cpuInfo.UsedPercent = gvar.New(percent[0]).String()
	return cpuInfo
}

// GetLoadInfo 负载信息
func (*sRuntime) GetLoadInfo() Load {
	info, err := load.Avg()
	if err != nil {
		fmt.Println("load.Avg() failed: ", err)
	}
	var loadInfo Load

	loadInfo.Load1 = gvar.New(info.Load1).String()
	loadInfo.Load5 = gvar.New(info.Load5).String()
	loadInfo.Load15 = gvar.New(info.Load15).String()
	return loadInfo
}

// GetMemInfo 内存信息
func (*sRuntime) GetMemInfo() Mem {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("mem.VirtualMemory() failed: ", err)
	}
	var memIn Mem
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
func (*sRuntime) GetHostInfo() Host {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("host.Info() failed: ", err)
	}
	var hostIn Host
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
func (*sRuntime) GetDiskInfo() Disk {
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
	var disk Disk
	disk.All = gvar.New(state.All / diskstate.MB).String()
	disk.Free = gvar.New(state.Free / diskstate.MB).String()
	disk.Available = gvar.New(state.Available / diskstate.MB).String()
	disk.Used = gvar.New(state.Used / diskstate.MB).String()
	disk.Usage = gvar.New(100 * state.Used / state.All).String()
	return disk
}

// GetNetInfo 网络信息
func (*sRuntime) GetNetInfo() Net {
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

	var kbsSentCacheKey = Util().ProjectName() + ":net_info:kbs_sent"
	var kbsRecvCacheKey = Util().ProjectName() + ":net_info:kbs_recv"
	var kbsTimeCacheKey = Util().ProjectName() + ":net_info:kbs_time"
	conn, err := g.Redis().Conn(Ctx)
	kbsSentCached, _ := conn.Do(Ctx, "GET", kbsSentCacheKey)
	kbsRecvCached, _ := conn.Do(Ctx, "GET", kbsRecvCacheKey)
	kbsTimeCached, _ := conn.Do(Ctx, "GET", kbsTimeCacheKey)

	//fmt.Println("kbsSentCached", kbsSentCached)
	//fmt.Println("kbsRecvCached", kbsRecvCached)
	//fmt.Println("kbsTimeCached", kbsTimeCached)

	conn.Do(Ctx, "SET", kbsSentCacheKey, kbsSent)
	conn.Do(Ctx, "SET", kbsRecvCacheKey, kbsRecv)
	conn.Do(Ctx, "SET", kbsTimeCacheKey, time.Now().Unix())
	var netInfo Net
	var seconds = gvar.New(time.Now().Unix()).Int() - kbsTimeCached.Int()
	if seconds > 0 {
		netInfo.KbsSent = (kbsSent - kbsSentCached.Int()) / seconds
		netInfo.KbsRecv = (kbsRecv - kbsRecvCached.Int()) / seconds
	} else {
		netInfo.KbsSent = kbsSent
		netInfo.KbsRecv = kbsRecv
	}
	defer conn.Close(Ctx)
	return netInfo
}

// SetServerStartAt 设置服务启动时间
func (*sRuntime) SetServerStartAt() bool {
	gcache.Set(Ctx, serverStartAtCacheKey, time.Now().Unix(), 0)
	return true
}

// GetServerStartAt 获取服务启动时间
func (*sRuntime) GetServerStartAt() *gvar.Var {
	get, _ := gcache.Get(Ctx, serverStartAtCacheKey)
	return get
}
