package main

import (
	"fmt"
	"time"
)

func recvFromUnbufferedChannel() {
	var ch = make(chan int, 3)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- 1
		ch <- 2
		ch <- 3

		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}

func recvFromNilChannel() {
	var ch chan int

	// 程序一直阻塞在这里
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	recvFromUnbufferedChannel()
	recvFromNilChannel()
}
