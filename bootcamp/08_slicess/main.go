package main

import (
	"fmt"
	"slices"
)

func main() {
	wel := "Welcome to slices"
	fmt.Println(wel)

	// a is a nil slice
	var a []byte

	// slice literal
	b := []byte{1, 2, 3, 4}

	// slice from an array
	c := b[1:3]
	fmt.Println("c:", c) // c: [2 3]

	// slice with make
	d := make([]byte, 1, 3)
	fmt.Println("d:", d) // d: [0]

	// slice with new
	e := *new([]byte)
	fmt.Println("e:", e) // e: []

	array := []int{1, 2, 3, 4, 5, 6}
	slice := array[1:3]

	// an integer array
	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	// create slice from an array
	s := numbers[4:7]
	fmt.Println(s) // [50 60 70]

	// ----------- append() - adds element to a slice
	primes := []int{2, 3}
	primes = append(primes, 5, 7)
	fmt.Println("Prime Numbers:", primes) // Prime Numbers: [2 3 5 7]

	// combine slices
	evens := []int{2, 4}
	odds := []int{1, 3}
	evens = append(evens, odds...)
	fmt.Println("Numbers:", evens) // Numbers: [2 4 1 3]

	// ----------- copy() - copy elements of one slice to another
	y := []byte{1, 2, 3, 4}
	x := []byte{5, 6, 7}
	copy(x, y)
	fmt.Println("Copy:", x) // Copy: [1 2 3] - because the size of x is 3 and it can only hold 3 elements

	// ----------- Equal() - compares two slices
	t1 := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t1, t2) {
		fmt.Println("Both are Equal: t1 == t2")
	}

	// ----------- len() - find the length of a slice
	fmt.Println("Length: ", len(a))            // Length:  0
	fmt.Println(slice, len(slice), cap(slice)) // [2 3] 2 5

	nums := []int{2, 4, 6, 8, 10}
	for i := 0; i < len(nums); i++ {
		fmt.Println("Num:", nums[i])
	}
}
