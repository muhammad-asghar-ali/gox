package cli

import (
	"crypton/crypton"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
)

type (
	CommandLine struct{}
)

func (cli CommandLine) Cmd() {
	fmt.Println("Usage: ")
	fmt.Println("get-balance -address ADDRESS - get the balance for that address")
	fmt.Println("create -address ADDRESS - create the cryption blockchain")
	fmt.Println("print - Prints the blocks in the chain")
	fmt.Println("send -from FROM -to TO -amount AMOUNT - send amount")
}

func (cli CommandLine) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.Cmd()
		runtime.Goexit()
	}
}

func (cli CommandLine) Print() {
	cytp := crypton.Crypton{}
	c := cytp.Contiune("")
	defer c.Database.Close()

	iter := c.Iterator()
	for {
		block := iter.Next()

		fmt.Printf("PrevHash: %x\n", block.PrevHash)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := crypton.NewProof(block)
		fmt.Printf("POW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func (cli CommandLine) Create(addr string) {
	cytp := crypton.Crypton{}
	c := cytp.Init("")
	defer c.Database.Close()

	fmt.Println("Finished!")
}

func (cli CommandLine) Balance(addr string) {
	cytp := crypton.Crypton{}
	c := cytp.Init("")
	defer c.Database.Close()

	balance := 0
	utxos := c.FindUTXO(addr)

	for _, out := range utxos {
		balance += out.Value
	}

	fmt.Printf("Balance of %s: %d\n", addr, balance)
}

func (cli CommandLine) Send(from, to string, amount int) {
	cytp := crypton.Crypton{}
	c := cytp.Contiune("")
	defer c.Database.Close()

	tx := crypton.NewTransection(from, to, amount, c)
	c.AddBlock([]*crypton.Transection{tx})

	fmt.Println("Success!")
}

func (cli CommandLine) Run() {
	cli.ValidateArgs()

	balance_cmd := flag.NewFlagSet("get-balance", flag.ExitOnError)
	cryption_cmd := flag.NewFlagSet("add", flag.ExitOnError)
	send_cmd := flag.NewFlagSet("send", flag.ExitOnError)
	print_cmd := flag.NewFlagSet("print", flag.ExitOnError)

	balance_addr := balance_cmd.String("address", "", "The address to get balance for")
	cryption_addr := cryption_cmd.String("address", "", "The address to send genesis block reward to")
	from := send_cmd.String("from", "", "Source wallet address")
	to := send_cmd.String("to", "", "Destination wallet address")
	amount := send_cmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "get-balance":
		err := balance_cmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "add":
		err := cryption_cmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := print_cmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := send_cmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.Print()
		runtime.Goexit()
	}

	if balance_cmd.Parsed() {
		if *balance_addr == "" {
			balance_cmd.Usage()
			runtime.Goexit()
		}
		cli.Balance(*balance_addr)
	}

	if cryption_cmd.Parsed() {
		if *cryption_addr == "" {
			cryption_cmd.Usage()
			runtime.Goexit()
		}
		cli.Create(*cryption_addr)
	}

	if print_cmd.Parsed() {
		cli.Print()
	}

	if send_cmd.Parsed() {
		if *from == "" || *to == "" || *amount <= 0 {
			send_cmd.Usage()
			runtime.Goexit()
		}

		cli.Send(*from, *to, *amount)
	}
}
