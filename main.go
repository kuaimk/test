package main

import (
	_ "aaa/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"aaa/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
