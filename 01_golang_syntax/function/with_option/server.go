package main

import "sync"

func main() {
	ch := make(chan struct{}) // 仅使用一个 channel 控制同步

	var wg sync.WaitGroup
	// 打印奇数

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 99; i += 2 {
			println("Odd:", i) // 打印奇数
			ch <- struct{}{}   // 通知偶数协程可以执行
			<-ch               // 等待偶数协程完成
		}
	}()

	// 打印偶数
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			<-ch                // 等待奇数协程通知
			println("Even:", i) // 打印偶数
			ch <- struct{}{}    // 通知奇数协程继续
		}
	}()

	// 防止主线程退出（可选：可以用 sync.WaitGroup 替代）
	wg.Wait()
}
