package main

import (
	_ "gf_cms/internal/packed"

	"gf_cms/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
