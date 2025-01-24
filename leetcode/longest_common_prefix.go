//go:build longest_common_prefix

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min := len(strs[0])
	for _, str := range strs {
		if len(str) < min {
			min = len(str)
		}
	}

	fmt.Println(min)
	for i := 0; i < min; i++ {
		char := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:min]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	r := longestCommonPrefix(strs)
	fmt.Println("rsult:", r)

	strs1 := []string{"dog", "racecar", "car"}
	r1 := longestCommonPrefix(strs1)
	fmt.Println("rsult:", r1)
}
