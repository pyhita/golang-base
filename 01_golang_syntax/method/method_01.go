package main

// go 方法案例一：说明如果需要修改类型 T 自身的值
// 应该使用的 *T 而不是 T

type T struct {
	name string
}

func (t T) M1(s string) {
	t.name = s
}

func (t *T) M2(s string) {
	t.name = s
}

func main() {
	var t T // t.name = ""
	print(t.name)

	t.M1("hello")
	print(t.name)

	t.M2("world")
	print(t.name)
}
