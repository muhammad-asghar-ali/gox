package main

import (
	"os"

	"crypton/cmd/cli"
)

func main() {
	defer os.Exit(0)

	cli := cli.CommandLine{}
	cli.Run()
}
