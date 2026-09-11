package main

import "fmt"

// Define a new type MySlice based on []int
type MySlice []int

// Constraint requires types whose core type is []int (using ~[]int)
type IntSliceConstraint interface {
	~[]int
}

func printFirstElement[S IntSliceConstraint](s S) int {
	// We can use indexing because the core type is a slice.
	return s[0]
}

func main() {
	ms := MySlice{10, 20, 30}
	fmt.Println(printFirstElement(ms)) // Prints: 10
	//ss := []string{"1", "2", "3"}
	//fmt.Println(printFirstElement(ss))
}

// The constraint IntSliceConstraint requires ~[]int, which means any type whose core type is []int.
// The core type of MySlice is its underlying type []int.It lets the compiler know that MySlice behaves like a slice.
