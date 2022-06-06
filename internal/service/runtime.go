package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

type sRuntime struct{}

var (
	insRuntime = sRuntime{}
)

func Runtime() *sRuntime {
	return &insRuntime
}

// GetAll 获取所有信息
func (*sRuntime) GetAll() {
	fmt.Println("==========CPU 信息==========")
	var cpuInfo = Runtime().GetCpuInfo()
	fmt.Println(cpuInfo)

	fmt.Println("\n==========负载 信息==========")
	var loadInfo = Runtime().GetLoadInfo()
	fmt.Println(loadInfo)

	fmt.Println("\n==========内存 信息==========")
	var memInfo = Runtime().GetMemInfo()
	fmt.Println(memInfo)

	fmt.Println("\n==========主机 信息==========")
	var hostInfo = Runtime().GetHostInfo()
	fmt.Println(hostInfo)

	fmt.Println("\n==========磁盘 信息==========")
	var diskInfo = Runtime().GetDiskInfo()
	fmt.Println(diskInfo)

	return
}

// GetCpuInfo CPU信息
func (*sRuntime) GetCpuInfo() []float64 {
	//cpuInfos, err := cpu.Info()
	//if err != nil {
	//	fmt.Println("get cpu info failed: ", err)
	//	return
	//}
	//
	//for _, ci := range cpuInfos {
	//	fmt.Println(ci)
	//}
	percent, _ := cpu.Percent(time.Second, false) // false表示CPU总使用率，true为单核
	return percent
}

// GetLoadInfo 负载信息
func (*sRuntime) GetLoadInfo() *load.AvgStat {
	info, err := load.Avg()
	if err != nil {
		fmt.Println("load.Avg() failed: ", err)
		return nil
	}
	return info
}

// GetMemInfo 内存信息
func (*sRuntime) GetMemInfo() *mem.VirtualMemoryStat {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("mem.VirtualMemory() failed: ", err)
		return nil
	}
	return memInfo
}

// GetHostInfo 主机信息
func (*sRuntime) GetHostInfo() *host.InfoStat {
	hostInfo, err := host.Info()
	if err != nil {
		fmt.Println("host.Info() failed: ", err)
		return nil
	}

	return hostInfo
}

// GetDiskInfo 磁盘信息
func (*sRuntime) GetDiskInfo() *disk.UsageStat {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("get disk partitions failed: ", err)
	}
	for _, part := range parts {
		partInfo, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Println("get part stat failed: ", err)
			return nil
		}
		return partInfo
	}
	return nil
}

// GetNetInfo 网络信息
func (*sRuntime) GetNetInfo() {
	netIOs, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("get net io counters failed: ", err)
		return
	}

	for _, netIO := range netIOs {
		fmt.Println(netIO) // 打印每张网卡信息
	}
}
