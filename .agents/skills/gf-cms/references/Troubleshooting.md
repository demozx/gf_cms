# 常见问题

## 1) gf-cli 未安装

- `make cli` 会通过 `wget` 下载并安装 gf。
- `make cli.install` 会检测 `gf -v`，未安装则自动安装。

## 2) 生成代码被手改导致冲突

- 如果你修改了 `internal/dao/**`、`internal/model/{do,entity}/**`、`internal/service/*.go`，下次 `gf gen dao` 很可能覆盖。
- 推荐做法：把业务逻辑放在 `internal/logic/**`、`internal/service/**` 的实现里，不改生成文件。

## 3) build.sh 与 upload 目录

- `build.sh` 会移动 `resource/public/upload/*` 到临时目录再执行 `gf build`。
- 如果 upload 目录不存在或为空，脚本可能失败；需要先创建目录或调整脚本。

## 4) Makefile 的 start/port

- `make start` 调用 `make port`，但仓库未提供 `port` 目标。
- 需要手动执行 `kubectl port-forward` 或补齐 Makefile。
