package main

import "fmt"

// defer 示例五：defer 后面的表达式是在被压入栈的时候求值的

func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(x int) {
			fmt.Println(x)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func main() {
	foo1()
	foo2()
	foo3()
}
