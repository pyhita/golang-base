package main

import "sync"

// channel03：channel 替代互斥锁

type Counter struct {
	value int
	ch    chan int
}

func NewCounter() *Counter {
	cter := &Counter{
		ch: make(chan int),
	}

	go func() {
		for {
			cter.value++
			cter.ch <- cter.value
		}
	}()
	return cter
}

func (c *Counter) Inc() int {
	return <-c.ch
}

func main() {
	cter := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println("goroutine ", i, " value ", cter.Inc())
		}(i + 1)
	}

	wg.Wait()
}
