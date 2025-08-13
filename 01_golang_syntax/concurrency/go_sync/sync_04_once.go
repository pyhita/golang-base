package main

import (
	"fmt"
	"sync"
)

type foo struct{}

var (
	once     sync.Once
	instance *foo
)

func GetInstance(id int) *foo {
	fmt.Println("goroutine ", id, " is getting instance")

	once.Do(func() {
		fmt.Println("goroutine ", id, " got instance")
		instance = &foo{}
	})

	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			GetInstance(i)
		}(i + 1)
	}

	wg.Wait()
	instance := GetInstance(1)
	fmt.Println("GetInstance: ", instance)
}
