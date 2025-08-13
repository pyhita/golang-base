package main

// 方法集合示例一：
// - T类型方法集合只包含所有receiver 参数是T的方法
// - *T类型方法集合包含所有receiver 参数是T和* T的方法

type I1 interface {
	M1()
	M2()
}

type A struct {
}

func (a A) M1() {

}

func (a *A) M2() {

}

func main() {
	var i I1
	a := A{}

	// T 类型方法集合只包含M1 所以没有实现接口I1
	i = a
	// *T 类型方法集合包含M1和M2 所以实现了接口I1
	i = &a
}
