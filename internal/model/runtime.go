package model

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
