package main

import (
	"fmt"
	"sync"
	"time"
)

// channel02：无缓冲channel 一对多的通知

type signal struct{}

func worker2(i int) {
	fmt.Printf("Worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}

func spawnGroup1(w func(int), n int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			<-groupSignal
			w(n + 1)
		}(i)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()

	return c
}

func main() {
	groupSignal := make(chan signal)
	done := spawnGroup1(worker2, 8, groupSignal)

	time.Sleep(5 * time.Second)
	close(groupSignal)

	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	select {
	case <-done:
		println("main exited")
	case <-timer.C:
		println("timed out")
	}
}
