package main

// defer 实例四：defer 可以修改函数的具名返回值

func baz(x int) (a, b int) {
	defer func() {
		a = a * 5
		b = b * 6
	}()

	a = x + 1
	b = x + 2
	return
}

func main() {
	println(baz(3))
}
