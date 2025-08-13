package main

import (
	"fmt"
	"time"
)

func demo1() {
	var nums = []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		go func() {
			time.Sleep(time.Second)
			fmt.Println(i, " ", num)
		}()
	}

	time.Sleep(time.Second * 6)
}

func demo2() {
	var nums = []int{1, 2, 3, 4, 5}
	for i, num := range nums {
		time.Sleep(time.Second)
		go func(i, v int) {
			fmt.Println(i, " ", v)
		}(i, num)
	}

	time.Sleep(time.Second * 6)
}

func main() {
	demo1()
	demo2()
}
