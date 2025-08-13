package main

import "fmt"

// 函数也可以绑定方法并且实现接口

// AddFunc MyAdd函数赋值给
// BinaryAdder接口。直接赋值是不行的，我们需要一个底层函数类型与
// MyAdd一致的自定义类型的显式转换，这个自定义类型就是
// MyAdderFunc，该类型实现了BinaryAdder接口，这样在经过
// MyAdderFunc的显式类型转换后，MyAdd被赋值给了BinaryAdder的变
// 量i。这样，通过i调用的Add方法实质上就是MyAdd函数。
type AddFunc interface {
	Add(a int, b int) int
}

type MyAddFunc func(int, int) int

func (m MyAddFunc) Add(a, b int) int {
	return m(a, b)
}

// MyAdd 函数相当于就是用户自定义的逻辑
// 不管用户定义的函数是哈，我们都希望用统一的接口变量接受
func MyAdd(a, b int) int {
	return a + b
}

func main() {
	var f AddFunc = MyAddFunc(MyAdd)
	fmt.Println(f.Add(1, 2))
}
