package main

import "fmt"

// 给定一个slice，去除切片中的"" 元素

func nonempty(nums []string) []string {
	i := 0
	for _, s := range nums {
		if s != "" {
			nums[i] = s
			i++
		}
	}

	return nums[:i]
}

func nonempty2(nums []string) []string {
	z := nums[:0]

	for _, s := range nums {
		if s != "" {
			z = append(z, s)
		}
	}

	return z
}

func main() {
	fmt.Println(nonempty([]string{"a", "b", "", "c"}))
	fmt.Println(nonempty2([]string{"a", "b", "", "c"}))
}
