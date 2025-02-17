package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	wel := "Welcome to the conversions"
	fmt.Println(wel)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the rating for pizza between 1 to 5:")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("thanks for the rating:", input)

	to_num, err := strconv.ParseInt(strings.TrimSpace(input), 32, 64)
	if err != nil {
		panic(err)
	}

	to_num += 1
	fmt.Println("plus 1 rating: ", to_num)

	to_flt, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}

	to_flt += 1
	fmt.Println("plus 1 rating: ", to_num)
}
