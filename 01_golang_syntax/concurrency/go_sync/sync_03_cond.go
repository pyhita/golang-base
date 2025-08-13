package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false

func worker(i int) {
	fmt.Println("worker ", i, " is working")
	time.Sleep(1 * time.Second)
	fmt.Println("worker ", i, " done")
}

func spawnGroup(f func(i int), n int, groupSignal *sync.Cond) <-chan struct{} {
	var wg sync.WaitGroup
	ch := make(chan struct{})

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			groupSignal.L.Lock()
			for !ready {
				// 条件不满足，暂时挂起，等待被唤醒
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Println("worker ", i, " start to work")
			f(i)
		}(i + 1)
	}

	go func() {
		wg.Wait()
		ch <- struct{}{}
	}()

	return ch
}

func main() {
	fmt.Println("start a group workers")

	groupSignal := sync.NewCond(&sync.Mutex{})
	done := spawnGroup(worker, 5, groupSignal)
	// 模拟处理ready
	time.Sleep(5 * time.Second)
	groupSignal.L.Lock()
	ready = true
	// 唤醒阻塞协程
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-done
	fmt.Println("the group of workers work done!")
}
