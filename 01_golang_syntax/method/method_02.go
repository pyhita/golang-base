package main

// go 方法示例二：不论方法的receiver 是啥都可以互相调用
// 因为golang 编译器会帮我们进行转换

type T struct {
	a int
}

func (t T) M1() {
}

func (t *T) M2() {
	t.a = 11
}

func main() {
	var t T
	t.M1()
	t.M2() // <=> (&t).M2()

	var pt = &T{}
	pt.M1() // <=> (*pt).M1()
	pt.M2()
}
