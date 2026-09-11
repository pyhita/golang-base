package main

import (
	"fmt"
	"time"
)

// Process 5 array elements using MAX 3 goroutines.

const maxConcurrency = 3

func worker(n int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		// calculate
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("%d process %d \n", n, j)
		res <- j * j
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	jobs := make(chan int, len(nums))
	res := make(chan int, len(nums))

	for _, n := range nums {
		jobs <- n
	}

	for i := 0; i < maxConcurrency; i++ {
		go worker(i+1, jobs, res)
	}

	// process res
	for i := 0; i < len(nums); i++ {
		fmt.Printf("processed value: %d \n", <-res)
	}
	close(jobs)
}
