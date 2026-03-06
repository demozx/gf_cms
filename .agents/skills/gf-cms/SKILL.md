---
name: gf-cms
description: gf_cms 项目专属技能集（本仓库）。聚焦本项目的目录结构、开发/构建/发布工作流、GoFrame CLI 生成规范，以及常见改造点（新增接口/路由/业务逻辑）。
license: MIT
---

# 重要规范

## 生成代码禁止手改

- `internal/service/*.go`、`internal/dao/**`、`internal/model/{do,entity}/**` 等带有 "Code generated and maintained by GoFrame CLI tool" 注释的文件，默认视为由 `gf` 工具生成维护：不要手动修改。
- 需要变更数据结构/表字段时：优先修改数据库结构/生成配置，然后运行 `gf gen dao` 重新生成。

## 本项目代码分层约定（按仓库现状）

- 入口：`main.go` -> `internal/cmd/cmd.go`。
- 路由：`internal/router/*`（按 backend/pc/mobile + api/view 分组）。
- 业务实现：`internal/logic/*`（每个模块在 `init()` 中 `service.RegisterXxx(New())` 注册到 `internal/service/*` 的接口）。
- 控制器：`internal/controller/**`（调用 `service.*()` 接口）。

## 配置与资源

- 本地运行需要 `config.yaml`；示例模板在 `manifest/config/config.yaml.example`。
- `gf build`/`gf docker` 的行为受 `config.yaml` 里的 `gfcli.build`、`gfcli.docker` 配置影响。

# 使用指南

- 项目技能文档索引：`./references/README.MD`
- 项目操作示例索引：`./examples/README.MD`
