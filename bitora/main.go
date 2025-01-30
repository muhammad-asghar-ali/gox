package main

import (
	"fmt"
	"strconv"

	"github.com/muhammad-asghar-ali/gox/bitora/core"
)

func main() {
	bitora := core.NewBitora()

	bitora.AddBlock("Send 1 BTC to John")
	bitora.AddBlock("Send 2 more BTC to John")

	for _, block := range bitora.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := core.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
