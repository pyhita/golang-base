package main

import "fmt"

// 用 slice 来模拟 stack
var stack []int

// push ele
func push(n int) {
	stack = append(stack, n)
}

// pop ele
func pop() int {
	if len(stack) == 0 {
		panic("stack is empty")
	}
	x := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return x
}

// top ele
func top() int {
	if len(stack) == 0 {
		panic("stack is empty")
	}
	return stack[len(stack)-1]
}

func main() {
	push(1)
	push(2)
	push(3)
	fmt.Println(stack)

	fmt.Println(top())

	fmt.Println(pop())
	fmt.Println(pop())
	fmt.Println(stack)
}
