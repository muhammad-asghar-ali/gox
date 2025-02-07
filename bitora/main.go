package main

import (
	"github.com/muhammad-asghar-ali/gox/bitora/cmd"
	"github.com/muhammad-asghar-ali/gox/bitora/core"
)

func main() {
	bitora := core.NewBitora()
	defer bitora.DB.Close()

	cli := cmd.CLI{Bit: bitora}
	cli.Run()
}
