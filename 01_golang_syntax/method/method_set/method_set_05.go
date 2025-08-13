package main

import "fmt"

// 方法集合示例五：
// - 如果某个结构体嵌入的多个接口含有相同的方法，此时编译不会通过，因为go不知道调用那个方法
// - 此时可以是该结构体重新实现重名的方法，这样实际就会调用重写的方法

type I3 interface {
	M1()
	M2()
	M3()
}

type I4 interface {
	M1()
	M2()
	M4()
}

type B struct {
	I3
	I4
}

func (b B) M1() {
	fmt.Println("B.M1()")
}

func main() {
	b := B{}
	b.M1() // 编译错误，因为不知道调用哪个接口的方法
}
