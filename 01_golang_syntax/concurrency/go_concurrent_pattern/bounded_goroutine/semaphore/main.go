package main

import (
	"fmt"
	"time"
)

// Process 5 array elements using MAX 3 goroutines.

const maxConcurrency = 3

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := make(chan int, len(nums))
	sem := make(chan struct{}, maxConcurrency)

	for i := 0; i < len(nums); i++ {
		// acquire lock
		sem <- struct{}{}
		go func(j int) {
			defer func() { <-sem }()
			time.Sleep(500 * time.Millisecond)
			res <- j * j
		}(nums[i])
	}

	// process res
	for i := 0; i < len(nums); i++ {
		fmt.Printf("processed value: %d \n", <-res)
	}
}
