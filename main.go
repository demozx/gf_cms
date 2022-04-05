package main

import (
	"gf_cms/internal/cmd"
	_ "gf_cms/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
