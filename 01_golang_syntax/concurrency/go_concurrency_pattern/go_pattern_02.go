package main

import (
	"errors"
	"fmt"
	"time"
)

// pattern2：父 goroutine 需要获取子 goroutine 的退出状态

var OK = errors.New("ok")

func worker2(args ...interface{}) error {
	if len(args) == 0 {
		return errors.New("invalid args")
	}
	interval, ok := args[0].(int)
	if !ok {
		return errors.New("invalid interval arg")
	}
	time.Sleep(time.Second * (time.Duration(interval)))
	return OK
}

func spawn2(f func(args ...interface{}) error, args ...interface{}) chan error {
	c := make(chan error)

	go func() {
		c <- f(args...)
	}()

	return c
}

func main() {
	c := spawn2(worker2, 3)
	err := <-c
	fmt.Printf("worker1 done: %v\n", err)

	c = spawn2(worker2)
	err = <-c
	fmt.Printf("worker2 done: %v\n", err)
}
