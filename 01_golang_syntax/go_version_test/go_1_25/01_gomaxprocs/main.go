package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Logical CPUs available:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
}

// Output:
// Logical CPUs available: 16
// GOMAXPROCS: 16
