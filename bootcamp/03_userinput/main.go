package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wel := "Welcome to the user input"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the rating for pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating", input)
	fmt.Printf("type of the input is %T\n", input)
}
