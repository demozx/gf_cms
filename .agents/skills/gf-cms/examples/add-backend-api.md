# 示例：新增一个后台 API

目标：新增 `POST /<backendPrefix>_api/ping` 返回服务端时间。

1) 在 `api/backendApi/` 新建 `ping.go`，定义请求/响应：

```go
package backendApi

import (
    "time"

    "github.com/gogf/gf/v2/frame/g"
)

type PingReq struct {
    g.Meta `tags:"BackendApi" method:"post" summary:"健康检查"`
}

type PingRes struct {
    Now time.Time `json:"now"`
}
```

2) 在 `internal/controller/backendApi/` 新建 `ping.go`，实现处理函数，内部调用 service（如果不需要复杂逻辑也可以直接返回）。

3) 在 `internal/router/backendApi.go` 的已登录分组里加入路由映射：

```go
"/ping": backendApi.Ping.Ping,
```

说明：该仓库大量路由使用 `group.ALLMap(g.Map{ ... })` 形式集中声明。
