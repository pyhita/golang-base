package main

import "fmt"

// defer 示例三：defer 可以拦截panic 并恢复

func bar() {
	fmt.Println("raise a panic!")
	panic("panic in bar")
}

func foo() {
	defer func() {
		// 拦截panic
		if err := recover(); err != nil {
			fmt.Println("recover in foo:", err)
		}
	}()
	bar()
}

func main() {
	foo()
}
