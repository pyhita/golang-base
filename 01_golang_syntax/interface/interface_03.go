package main

import "fmt"

// 接口3：装箱之后的数据存在于新的内存空间，和原数据不在存在联系
func main() {
	var n int = 61
	var ei interface{} = n
	n = 62
	fmt.Println("data in box:", ei)

	var m int = 51
	ei = &m
	m = 52

	p := ei.(*int)
	fmt.Println("data in box:", *p)
}
