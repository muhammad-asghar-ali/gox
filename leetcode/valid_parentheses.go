//go:build valid_parentheses

package main

import (
	"fmt"
)

func isValid(s string) bool {
	var stack []string

	for i := 0; i < len(s); i++ {
		cur := string(s[i])

		n := len(stack)
		if n > 0 {
			last := stack[n-1]
			if isPair(last, cur) {
				stack = stack[:n-1]
				continue
			}
		}

		stack = append(stack, cur)
	}

	return len(stack) == 0
}

func isPair(last, cur string) bool {
	if (last == "(" && cur == ")") ||
		(last == "{" && cur == "}") ||
		(last == "[" && cur == "]") {
		return true
	}

	return false
}

func main() {
	s := "([]){[}]}"
	res := isValid(s)
	fmt.Println("result: ", res)

	s1 := "([]){[]}"
	res1 := isValid(s1)
	fmt.Println("result1: ", res1)

	s2 := "([])"
	res2 := isValid(s2)
	fmt.Println("result2: ", res2)
}
