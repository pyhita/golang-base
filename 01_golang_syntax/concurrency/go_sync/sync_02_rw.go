package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu     sync.RWMutex
	values map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock() // 写锁
	defer c.mu.Unlock()
	c.values[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mu.RLock() // 读锁
	defer c.mu.RUnlock()
	return c.values[key]
}

func main() {
	c := SafeCounter{values: make(map[string]int)}

	// 启动多个goroutine进行写操作
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				c.Inc("somekey")
			}
		}()
	}

	// 启动多个goroutine进行读操作
	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				fmt.Println(c.Value("somekey"))
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Final value:", c.Value("somekey"))
}
