package backendApi

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetRuntimeInfoApiReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"获取运行时信息"`
}
type GetRuntimeInfoApiRes struct {
	Load string `json:"load" "dc": "负载"`
	Cpu  string `json:"cpu" "dc": "CPU"`
	Mem  string `json:"mem" "dc": "内存"`
	Disk string `json:"disk" "dc": "磁盘"`
}
