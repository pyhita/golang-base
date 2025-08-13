package main

import "fmt"

// 接口1：接口定义 和 类型断言操作
type MyInterface interface {
	M1()
}

type T int

func (T) M1() {
	println("T's M1")
}

func main() {
	var t T
	var i interface{} = t
	v1, ok := i.(MyInterface)
	if !ok {
		panic("the value of i is not MyInterface")
	}
	v1.M1()
	fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T

	i = int64(13)
	v2, ok := i.(MyInterface)
	fmt.Printf("the type of v2 is %T\n", v2) // the type of v2 is <nil>
	// v2 = 13 //  cannot use 1 (type int) as type MyInterface in assignment: int does not implement MyInterface (missing M1   method)
}
