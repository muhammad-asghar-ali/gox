package main

import (
	"fmt"
	"unsafe"
)

func main() {
	wel := "Welcome to arrays"
	fmt.Println(wel)

	f := [4]string{}

	f[0] = "a"
	f[1] = "b"
	f[3] = "c"

	for i := range f {
		fmt.Println(i, &f[i])
	}

	fmt.Println("array: ", f)                // [a b  c]
	fmt.Println("length of array: ", len(f)) // 4

	arr := [7]byte{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr: ", &arr) // &[1 2 3 4 5 6 7]

	// print the value with address
	for i := range arr {
		fmt.Println(i, &arr[i])
	}

	a := [3]int{99, 100, 101}
	p := unsafe.Pointer(&a[0])

	a1 := unsafe.Pointer(uintptr(p) + 8)
	a2 := unsafe.Pointer(uintptr(p) + 16)

	fmt.Println(*(*int)(p))  // 99
	fmt.Println(*(*int)(a1)) // 100
	fmt.Println(*(*int)(a2)) // 101

	// type
	b := [4]byte{}
	fmt.Printf("%T\n", b) // [4]uint8

	// ------------------------------------------------ Array literals
	var arr1 [10]int
	fmt.Println("arr1", arr1) // [0 0 0 0 0 0 0 0 0 0]

	// With value, infer-length
	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr2", arr2) // [1 2 3 4 5 6 7]

	// With index, infer-length
	arr3 := [...]int{11: 3}
	fmt.Println("arr3", arr3) // [0 0 0 0 0 0 0 0 0 0 0 3]

	// Combined index and value
	arr4 := [5]int{1, 4: 5} // [1 0 0 0 5]
	fmt.Println("arr4", arr4)

	arr5 := [5]int{2: 3, 4, 4: 5}
	fmt.Println("arr5", arr5) // [0 0 3 4 5]

	// ------------------------------------------------ Array operations
	c := [5]int{1, 2, 3}
	println(len(c)) // 5
	println(cap(c)) // 5

	s := [5]int{0, 1, 2, 3, 4}
	fmt.Println(s[1:3]) // [1 2]
	fmt.Println(s[:3])  // [0 1 2]
	fmt.Println(s[2:])  // [2 3 4]
}
