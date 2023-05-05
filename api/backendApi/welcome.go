package backendApi

import (
	"gf_cms/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRuntimeInfoApiReq struct {
	g.Meta `tags:"BackendApi" method:"post" summary:"获取运行时信息"`
}
type GetRuntimeInfoApiRes struct {
	Load                     model.Load `json:"load" dc:"负载"`
	LoadPercent              float32    `json:"loadPercent" dc:"负载"`
	Cpu                      model.Cpu  `json:"cpu" dc:"CPU"`
	CPUNum                   int        `json:"cpuNum" dc:"CPU核心数"`
	Mem                      model.Mem  `json:"mem" dc:"内存"`
	Disk                     model.Disk `json:"disk" dc:"磁盘"`
	Net                      model.Net  `json:"net" dc:"网络"`
	ServerStartDuration      string     `json:"serverStartDuration" dc:"服务运行时长"`
	NumGoroutine             int        `json:"numGoroutine" dc:"goroutine运行数量"`
	MysqlProcessNum          int        `json:"mysqlProcessNum" dc:"MySql进程数"`
	MySqlMaxConnectionsNum   int        `json:"mySqlMaxConnectionsNum" dc:"MySql最大连接数"`
	MySqlCurrConnectionsNum  int        `json:"mySqlCurrConnectionsNum" dc:"MySql当前连接数"`
	RedisMaxClientsNum       int        `json:"redisMaxClientsNum" dc:"Redis最大连接数"`
	RedisConnectedClientsNum int        `json:"redisConnectedClientsNum" dc:"Redis当前连接数"`
	Pid                      int        `json:"pid" dc:"服务Pid"`
}
