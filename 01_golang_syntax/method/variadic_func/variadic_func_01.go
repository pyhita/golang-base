package variadic_func

// 变长参数示例一：解释变长参数函数是啥，如何定义，如何传参

// 变长参数args 只能有一个，并且必须是最后一个参数
// 变长参数类型是切片，可以传入切片，也可以传入多个参数
func sum(args ...int) int {
	var total int

	for _, v := range args {
		total += v
	}

	return total
}

func main() {
	a, b, c := 1, 2, 3
	println(sum(a, b, c))
	nums := []int{4, 5, 6}
	println(sum(nums...))
}
