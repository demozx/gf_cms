# gf_cms 项目概览

## 入口与启动

- `main.go` 通过 `cmd.Main.Run(...)` 启动。
- `internal/cmd/cmd.go`:
  - 设置模板分隔符 `${ ... }$`。
  - 初始化 runtime 运行信息。
  - 根据缓存驱动选择 session 存储（memory/redis）。
  - 注册模板全局方法（`internal/logic/viewBindFun`）。
  - 注册路由（`internal/router/router.go`）。

## 路由组织

- 统一入口：`internal/router/router.go`。
- 分组文件：
  - `internal/router/backendApi.go` / `internal/router/backendView.go`
  - `internal/router/pcApi.go` / `internal/router/pcView.go`
  - `internal/router/mobileApi.go` / `internal/router/mobileView.go`

## 业务与依赖注入模式（当前项目的写法）

- `internal/service/*.go` 里定义接口 `IXxx` + `RegisterXxx` + `Xxx()` 单例获取。
- `internal/logic/<module>/*.go` 里实现具体逻辑，并在 `init()` 调 `service.RegisterXxx(New())`。
- `internal/controller/**` 调用 `service.Xxx()`，控制器自身尽量薄。

这一套的关键点：

- 每个模块必须在 `internal/logic/logic.go` 被 `_` import 才会触发 `init()` 注册。
- 新增模块时，除了新增 `logic/<module>` 之外，通常还需要更新 `internal/logic/logic.go` 的 import 列表（该文件标注为生成文件，通常由工具维护）。

## 配置文件

- 示例配置：`manifest/config/config.yaml.example`。
- 本地/运行时配置：`config.yaml`（生产部署 README 建议放在二进制同目录）。
- `gfcli.build`/`gfcli.docker` 配置用于 `gf build`、`gf docker`。
