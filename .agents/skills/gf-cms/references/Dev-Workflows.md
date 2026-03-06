# 开发工作流

本项目是 GoFrame v2 工程，README/Makefile 已经固化了一些常用命令。

## 0. 初始化（首次）

1) 导入数据库：`gfcms_demo.sql`

2) 生成配置文件：

```bash
cp manifest/config/config.yaml.example config.yaml
```

3) 按需修改 `config.yaml`：数据库/redis/域名绑定等。

## 1. 运行（开发）

- 推荐（安装 gf-cli 后）：

```bash
gf run main.go
```

- 不使用 gf-cli：

```bash
go build -o ./main main.go
./main
```

## 2. 代码生成（DAO/DO/Entity）

- 生成 DAO：

```bash
make dao
```

等价于：`gf gen dao`

注意：生成文件默认禁止手改（见 `internal/dao/**`、`internal/model/{do,entity}/**`、`internal/service/*.go` 的头注释）。

## 3. 构建（打包资源）

- 直接构建：

```bash
gf build
```

- 使用仓库脚本（会临时挪走上传目录，避免被打包/覆盖）：

```bash
./build.sh
```

`gf build` 的打包目录/输出目录/packDst 由 `config.yaml` 的 `gfcli.build` 配置控制。

## 4. Docker 镜像

- 构建镜像：

```bash
make image
```

- 构建并推送：

```bash
make image.push
```

说明：`make image` 内部使用 `gf docker`，镜像 tag 会根据 `git log` 生成；工作区 dirty 时会追加 `.dirty`。

## 5. Kubernetes 部署（kustomize + kubectl）

```bash
make deploy ENV=develop
```

说明：该目标会在 `temp/` 生成 kustomize 输出 YAML 并 apply，然后 patch deployment 标签触发滚动。

注意：Makefile 中 `make start` 会调用 `make port`，但当前仓库未提供 `port` 目标（需要自行补充 `kubectl port-forward`）。

## 工具可用性提示

- 本仓库环境里可能没有 `rg` (ripgrep)。写脚本/说明时，建议准备 `grep` 作为回退方案。
