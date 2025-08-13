package main

import (
	"fmt"
	"time"
)

// channel01： 无缓冲channel 可以用来一对一的通知

type signal struct{}

func worker1(i int) {
	fmt.Printf("Worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}

func spawnWorker1(w func(int)) <-chan signal {
	c := make(chan signal)
	go func() {
		w(0)
		c <- signal{}
	}()

	return c
}

func main() {
	done := spawnWorker1(worker1)
	<-done
	println("main exited")
}
