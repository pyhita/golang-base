package main

import "sync"

type CounterMutex struct {
	i int
	sync.Mutex
}

func NewCounterMutex() *CounterMutex {
	return &CounterMutex{i: 0}
}

func (c *CounterMutex) Inc() int {
	c.Lock()
	defer c.Unlock()
	c.i++
	return c.i
}

func main() {
	cter := NewCounterMutex()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			println("worker ", i, " value ", cter.Inc())
		}(i)
	}

	wg.Wait()
}
