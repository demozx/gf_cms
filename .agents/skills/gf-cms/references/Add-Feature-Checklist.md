# 新增模块/功能清单（按仓库惯例）

## 新增后台 API（backend_api）

1) 定义 API 请求/响应结构体（`api/backendApi/*.go`）

- 使用 `g.Meta` 声明 tags/method/summary。
- 使用 `v:"..."` 做参数校验。

2) 实现控制器（`internal/controller/backendApi/*.go`）

- 控制器只做参数解析/调用 service/统一返回。

3) 实现逻辑层并注册（`internal/logic/<module>/*.go`）

- `init()` 中 `service.RegisterXxx(New())`。
- 将数据库访问集中在 logic/service 层。

4) 路由注册（`internal/router/backendApi.go`）

- 按现有分组方式加入 `group.ALLMap(g.Map{ ... })`。

5) 如果是新增“服务接口”（`internal/service/*.go`）

- 该目录通常由 GoFrame CLI 生成维护；优先用工具生成，避免手写。

## 新增数据表/字段

1) 修改数据库结构
2) 运行 `gf gen dao`（或 `make dao`）
3) 检查生成代码与现有业务逻辑的字段映射/JSON 输出
