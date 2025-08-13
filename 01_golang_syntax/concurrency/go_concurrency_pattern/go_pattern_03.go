package main

import (
	"fmt"
	"sync"
	"time"
)

// pattern3：父 goroutine 需要等待子 goroutine 退出

func worker3(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	interval, ok := args[0].(int)
	if !ok {
		return
	}
	time.Sleep(time.Second * (time.Duration(interval)))
}

func spawnGroup(f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})

	wg := sync.WaitGroup{}
	//for _, i := range args {
	//	wg.Add(1)
	//	// i := i
	//	go func(n interface{}) {
	//		defer wg.Done()
	//		f(n)
	//		// fmt.Printf("worker %d done\n", i)
	//		v, ok := n.(int)
	//		if ok {
	//			println(v)
	//		}
	//	}(i)
	//}

	for i := 1; i <= len(args); i++ {
		wg.Add(1)
		go func(i int) {
			name := fmt.Sprintf("worker-%d:", i)
			f(args...)
			println(name, "done")
			wg.Done() // worker done!
		}(i)
	}

	go func() {
		wg.Wait()
		fmt.Println("all workers done")
		c <- struct{}{}
	}()

	return c
}

func main() {
	c := spawnGroup(worker3, 1, 2, 3, 4, 5)
	<-c

	fmt.Println("main goroutine done")
}
