package backendApi

import (
	"gf_cms/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

type GetRuntimeInfoApiReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"获取运行时信息"`
}
type GetRuntimeInfoApiRes struct {
	Load                service.Load `json:"load" "dc": "负载"`
	LoadPercent         float32      `json:"loadPercent" "dc": "负载"`
	Cpu                 service.Cpu  `json:"cpu" "dc": "CPU"`
	CPUNum              int          `json:"CPUNum" "dc": "CPU核心数"`
	Mem                 service.Mem  `json:"mem" "dc": "内存"`
	Disk                service.Disk `json:"disk" "dc": "磁盘"`
	Net                 service.Net  `json:"net" "dc": "网络"`
	ServerStartDuration string       `json:"serverStartDuration" "dc": "服务运行时长"`
}
