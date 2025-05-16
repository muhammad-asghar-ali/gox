package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/muhammad-asghar-ali/gox/bitora/core"
)

type CLI struct {
	Bit *core.Bitora
}

func (cli *CLI) Run() {
	cli.ValidateArgs()

	switch os.Args[1] {
	case "add":
		cli.handle_add(os.Args[2:])
	case "display":
		cli.Display()
	case "help":
		cli.DisplayUsage()
	default:
		fmt.Println("Error: Invalid command.")
		cli.DisplayUsage()
		os.Exit(1)
	}
}

func (cli *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.DisplayUsage()
		os.Exit(1)
	}
}

func (cli *CLI) handle_add(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addData := addCmd.String("data", "", "Block data")

	if err := addCmd.Parse(args); err != nil {
		fmt.Println("Error parsing arguments.")
		os.Exit(1)
	}

	if *addData == "" {
		fmt.Println("Error: Block data cannot be empty.")
		addCmd.Usage()
		os.Exit(1)
	}

	cli.AddBlock(*addData)
}

func (cli *CLI) AddBlock(data string) {
	cli.Bit.AddBlock(data)
	fmt.Println("Block successfully added!")
}

func (cli *CLI) Display() {
	bci := cli.Bit.Iterator()

	for {
		block := bci.Next()
		fmt.Println("-----------------------------------")
		fmt.Printf("Prev. Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW Valid: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println("-----------------------------------")

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli *CLI) DisplayUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add -data DATA  : Add a new block with DATA")
	fmt.Println("  display         : display the blockchain")
	fmt.Println("  help            : Show this help message")
}
