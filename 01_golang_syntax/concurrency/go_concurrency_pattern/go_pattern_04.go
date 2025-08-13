package main

import "time"

// pattern4: 父 goroutine 等待子 goroutine 有限时间，超过了时间，不再继续等待

func main() {
	done := spawnGroup(worker3, 1, 2, 3, 4, 5)
	println("spawn a group of workers")
	timer := time.NewTimer(time.Second * 3)

	select {
	case <-done:
		println("all workers done")
	case <-timer.C:
		println("timeout")
	}
}
