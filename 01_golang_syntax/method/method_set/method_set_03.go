package main

import (
	"fmt"
	"reflect"
)

// 方法集合示例三：结构体嵌入接口，不论是 T 还是 *T 的方法集合都包含所嵌入接口的方法集合

type Interface interface {
	M1()
	M2()
}

type T struct {
	Interface
}

func (T) M3() {}

func main() {
	DumpMethodSet((*Interface)(nil))
	var t T
	var pt *T
	DumpMethodSet(&t)
	DumpMethodSet(&pt)
}

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
