package main

import "fmt"

// 函数完全可以跟其他类型一样，当做一种类型进行使用
// 可以声明函数类型的数组、切片、channel、map等
type binaryCalcFunc func(int, int) int

func main() {
	var i interface{} = binaryCalcFunc(func(a int, b int) int {
		return a + b
	})

	v, ok := i.(binaryCalcFunc)
	if !ok {
		fmt.Println("type assert error")
	} else {
		fmt.Println(v(1, 2))
	}

	// 定义函数类型的channel
	ch := make(chan binaryCalcFunc, 10)
	ch <- func(a int, b int) int {
		return a + b
	}
	f := <-ch
	fmt.Println(f(1, 2))

	// 定义函数的slice
	funs := []binaryCalcFunc{
		func(a int, b int) int {
			return a + b
		},
		func(a int, b int) int {
			return a - b
		},
	}

	fmt.Println(funs[1](1, 2))

	// 定义函数类型的map
	m := map[string]binaryCalcFunc{
		"+": func(a int, b int) int {
			return a + b
		},
	}

	fmt.Println(m["+"](1, 2))
	// 函数不能进行比较，无法作为map的key
	//m2 := map[binaryCalcFunc]string{
	//	func(a int, b int) int {
	//}
}
