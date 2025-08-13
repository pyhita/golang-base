package main

// 方法集合示例七：结构体嵌入结构体
// - 类型T的方法集合 = T1的方法集合 + *T2的方法集合
// - 类型*T的方法集合 = *T1的方法集合 + *T2的方法集合

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type T1 T

func main() {
	var t T
	var pt *T
	var t1 T1
	var pt1 *T1

	DumpMethodSet(t)
	DumpMethodSet(t1)

	DumpMethodSet(pt)
	DumpMethodSet(pt1)
}
