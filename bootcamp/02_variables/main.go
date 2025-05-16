package main

import (
	"fmt"
)

const Token = "asdsdsadasds"
const TokenType string = "asdsdsadasds"

func main() {
	var username string = "username"
	fmt.Println(username)                        // username
	fmt.Printf("variable type : %T\n", username) // string

	var isActive bool = false
	fmt.Println(isActive)                           // false
	fmt.Printf("variable type is : %T\n", isActive) // bool

	var sm uint = 3
	fmt.Println(sm)                           // 3
	fmt.Printf("variable type is : %T\n", sm) // uint

	var ft float32 = 4.8888888
	fmt.Println(ft)                           // 4.888889
	fmt.Printf("variable type is : %T\n", ft) // float32

	// default values
	var dfInt int
	fmt.Println(dfInt)                           // 0
	fmt.Printf("variable type is : %T\n", dfInt) // int

	var dfFloat float64
	fmt.Println(dfFloat)                           // 0
	fmt.Printf("variable type is : %T\n", dfFloat) // float64

	var dfStr string
	fmt.Println(dfStr)                           //
	fmt.Printf("variable type is : %T\n", dfStr) // string

	var dfBl bool
	fmt.Println(dfBl)                           // false
	fmt.Printf("variable type is : %T\n", dfBl) // bool

	// implicit type
	var web = "youtube"
	fmt.Println(web)                           // youtube
	fmt.Printf("variable type is : %T\n", web) // string

	// no war style
	users := 100
	fmt.Println(users)                           // 100
	fmt.Printf("variable type is : %T\n", users) // int

	// const
	fmt.Println(Token)                           // asdsdsadasds
	fmt.Printf("variable type is : %T\n", Token) // string
}
