package main

import (
	"time"
)

// pattern5：通知并等待一个 goroutine 退出

func worker4(t int) {
	time.Sleep(time.Second * (time.Duration(t)))
}

func spawn4(f func(n int), n int) chan struct{} {
	done := make(chan struct{})
	job := make(chan int)

	go func() {
		select {
		// 接收到工作信号
		case <-job:
			f(n)
		// 收到退出信号
		case <-done:
			done <- struct{}{}
		}
	}()

	return done
}

func main() {
	quit := spawn4(worker4, 5)
	println("spawn a worker goroutine")
	time.Sleep(5 * time.Second)
	// 通知新创建的goroutine退出
	println("notify the worker to exit...")
	quit <- struct{}{}
	timer := time.NewTimer(time.Second * 10)
	defer timer.Stop()
	select {
	case <-quit:
		println("worker done")
	case <-timer.C:
		println("wait worker exit timeout")
	}
}
