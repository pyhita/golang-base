package main

import (
	"log"
	"sync"
	"time"
)

// channel04：channel 替代信号量
var jobs = make(chan int, 8)
var active = make(chan struct{}, 3)

func main() {

	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i
		}
		// 只是停止发送，已经在channel中的数据还可以读取
		close(jobs)
	}()

	wg := sync.WaitGroup{}
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
		}(j)
	}

	wg.Wait()
}
