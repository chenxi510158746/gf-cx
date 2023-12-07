package main

import (
	_ "gf-cx/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "gf-cx/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-cx/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
