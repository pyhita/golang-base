package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Go(func() {
		fmt.Println("Hello World 1")
		time.Sleep(1 * time.Second)
	})

	wg.Go(func() {
		fmt.Println("Hello World 2")
		time.Sleep(2 * time.Second)
	})

	wg.Wait()
	fmt.Println("main exit")
}
