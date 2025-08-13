package main

import (
	"fmt"
	"time"
)

/*
当countdown函数返回时，它会停止从tick中接收事件，但是ticker这个goroutine还依然存活，继续徒劳地尝试向channel中发送值，
然而这时候已经没有其它的goroutine会从该channel中接收值了——这被称为goroutine泄露

Tick函数挺方便，但是只有当程序整个生命周期都需要这个时间时我们使用它才比较合适。否则的话，我们应该使用下面的这种模式：

ticker := time.NewTicker(1 * time.Second)
<-ticker.C    // receive from the ticker's channel
ticker.Stop() // cause the ticker's goroutine to terminate

*/

func main() {
	fmt.Println("Commencing countdown.  Press return to abort.")

	abort := make(chan struct{})
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()

}

func launch() {
	fmt.Println("Lift off!")
}
