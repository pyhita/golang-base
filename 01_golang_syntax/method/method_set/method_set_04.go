package main

import "fmt"

// 方法集合示例四：结构体调用方法，先看自己的方法集合有没有该方法，如果没有继续查看所属字段的方法集合有没有改方法

type I2 interface {
	M1()
	M2()
}

type AA struct {
	I2
}

func (a AA) M1() {
	fmt.Println("AA's M1")
}

type SS struct {
}

func (s SS) M1() {
	fmt.Println("SS's M1")
}

func (s SS) M2() {
	fmt.Println("SS's M2")
}

func main() {
	s := SS{}
	aa := AA{
		s, // 这里s实现了I2接口，所以AA的方法集合中包含M1和M2
	}

	aa.M1() // AA's M1
	aa.M2() // SS's M2
}
