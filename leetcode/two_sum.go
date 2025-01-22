//go:build two_sum

package main

import "fmt"

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
*/
func twoSum(nums []int, target int) []int {
	// store numbers and their corresponding indices
	mp := make(map[int]int)

	for i, num := range nums {
		// difference between the target and the current number
		c := target - num

		// if the difference already exists in the map
		if idx, f := mp[c]; f {
			// return the indices of the current number and the number that adds up to the target
			return []int{idx, i}
		}

		// if it doesn't exist, add the current number and its index to the map
		mp[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 8, 11, 15}
	target := 9

	res := twoSum(nums, target)
	fmt.Println("result:", res)

	nums1 := []int{3, 2, 4}
	target1 := 6

	res1 := twoSum(nums1, target1)
	fmt.Println("result 1: ", res1)

	nums2 := []int{3, 3}
	target2 := 6

	res2 := twoSum(nums2, target2)
	fmt.Println("result 2: ", res2)
}
