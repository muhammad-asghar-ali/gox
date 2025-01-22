//go:build valid_parentheses

package main

import (
	"fmt"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/
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
