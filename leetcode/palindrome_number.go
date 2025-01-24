//go:build palindrome_number

package main

import "fmt"

// Given an integer x, return true if x is a palindrome, and false otherwise.
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original, reversed := x, 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

func main() {
	n1 := 121
	fmt.Println("result 1:", isPalindrome(n1))

	n2 := 1231
	fmt.Println("result 2:", isPalindrome(n2))

	n3 := 1221
	fmt.Println("result 3:", isPalindrome(n3))
}
