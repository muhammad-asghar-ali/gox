//go:build two_sum

package main

import "fmt"

func romanToInt(s string) int {
	var romanValues = [128]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)
	if n == 0 {
		return total
	}

	next := romanValues[s[n-1]]
	total += next

	for i := n - 2; i >= 0; i-- {
		curr := romanValues[s[i]]
		if curr < next {
			total -= curr
		} else {
			total += curr
		}

		next = curr
	}

	return total
}

func main() {
	s1 := "III"
	t1 := romanToInt(s1)
	fmt.Println("t1", t1)

	s2 := "LVIII"
	t2 := romanToInt(s2)
	fmt.Println("t2", t2)

	s3 := "MCMXCIV"
	t3 := romanToInt(s3)
	fmt.Println("t3", t3)
}
