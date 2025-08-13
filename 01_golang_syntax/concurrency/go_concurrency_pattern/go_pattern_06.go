package main

import (
	"sync"
	"time"
)

// pattern6：通知并等待多个 goroutine 退出

func worker5(t int) {
	time.Sleep(time.Second * (time.Duration(t)))
}

func spawnGroup2(f func(n int), _ int) chan struct{} {
	done := make(chan struct{})
	job := make(chan int)

	wg := sync.WaitGroup{}
	// 启动一组 goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			println("worker-", n)
			for {
				v, ok := <-job
				if !ok {
					print("done worker ", n, "\n")
					return
				}
				f(v)
			}
		}(i)
	}

	go func() {
		// 监听退出信号
		<-done
		// 发送退出信号
		close(job)
		// 等待所有 goroutine 退出
		wg.Wait()
		// 告诉 main 所有的worker 都已经退出
		done <- struct{}{}
	}()

	return done
}

func main() {
	println("spawn g group of goroutines")
	done := spawnGroup2(worker5, 1)

	time.Sleep(time.Second * 3)
	// 发送exit 信号
	done <- struct{}{}
	timer := time.NewTimer(time.Second * 3)
	select {
	case <-done:
		println("all workers done")
	case <-timer.C:
		println("timeout exit")
	}
}
