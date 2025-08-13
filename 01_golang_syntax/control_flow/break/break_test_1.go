package main

import (
	"fmt"
	"time"
)

func testBreak() {
	var ch = make(chan string)

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			case <-ch:
				fmt.Println("exiting")
				break
			}
		}
	}()

	time.Sleep(3 * time.Second)
	ch <- "exit"
	time.Sleep(3 * time.Second)
}

func main() {
	testBreak()
}
