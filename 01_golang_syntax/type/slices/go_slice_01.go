package main

import "fmt"

func main() {
	// The size specifies the length. The capacity of the slice is
	//	equal to its length. A second integer argument may be provided to
	//	specify a different capacity; it must be no smaller than the
	//	length. For example, make([]int, 0, 10) allocates an underlying array
	//	of size 10 and returns a slice of length 0 and capacity 10 that is
	//	backed by this underlying array.

	// 示例1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
}
