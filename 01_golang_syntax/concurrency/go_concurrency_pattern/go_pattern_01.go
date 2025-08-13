package main

import "time"

// pattern1：父 goroutine 需要等待子 goroutine 的退出，等待一个goroutine退出
func worker1(args ...interface{}) {
	if len(args) == 0 {
		return
	}
	interval, ok := args[0].(int)
	if !ok {
		return
	}
	time.Sleep(time.Second * (time.Duration(interval)))
}

func spawn1(f func(args ...interface{}), args ...interface{}) chan struct{} {
	c := make(chan struct{})

	go func() {
		// 执行业务逻辑
		f(args...)
		// 发送退出信号
		c <- struct{}{}
	}()

	return c
}

func main() {
	done := spawn1(worker1, 5)
	println("spawn a worker goroutine")
	<-done
	println("worker done")
}
