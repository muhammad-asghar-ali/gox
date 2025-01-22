package main

import (
	"fmt"
)

func main() {
	wel := "Welcome to pointer"
	fmt.Println(wel)

	// declaring Pointers
	var ptr *int
	fmt.Println("vlaue of ptr is:", ptr) // <nil>

	num := 44
	ptr1 := &num

	// using Pointers
	fmt.Println("value of ptr1 is:", ptr1) // 0x1400000e130

	// dereferencing Pointers
	fmt.Println("value of *ptr1 is:", *ptr1) // 44
	fmt.Println("value of &ptr1 is:", &ptr1) // 0x14000058038

	// changing Values via Pointers
	*ptr1 += 2
	fmt.Println("value of num is:", num)   // 46
	fmt.Println("value of &num is:", &num) // 0x1400000e130

	//
	x := 5
	change(x)
	fmt.Println("value remains the same: ", x) // 5

	change_ptr(&x)
	fmt.Println("updated value: ", x) // 10

	/*
		In Go, passing pointers to functions is often necessary and efficient,
		but the recommendation to avoid returning pointers from functions unless
		necessary is rooted in how memory management and garbage collection (GC) work in Go.
	*/
	// CreatePoint uses stack allocation
	p1 := CreatePoint(1, 2)

	// CreatePointPointer causes heap allocation
	p2 := CreatePointPointer(3, 4)

	fmt.Println(p1, p2)
}

func change(val int) {
	val = 10
}

func change_ptr(val *int) {
	*val = 10
}

type Point struct {
	X, Y int
}

// Return by value (stack allocation)
func CreatePoint(x, y int) Point {
	return Point{X: x, Y: y}
}

// Return by pointer (heap allocation)
func CreatePointPointer(x, y int) *Point {
	return &Point{X: x, Y: y}
}
