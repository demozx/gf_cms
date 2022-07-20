package main

import (
	_ "gf_cms/internal/packed"

	_ "gf_cms/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"gf_cms/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
