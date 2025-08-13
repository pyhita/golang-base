package main

import "fmt"

// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。

// 向左旋转 n 个字符
func rotateNLeft(nums []int, n int) []int {
	// 1 2 3 4 5
	// 3 4 5 1 2

	// s1: 1 2 3 4 5 1 2
	// s2: 3 4 5 1 2
	n %= len(nums)
	nums = append(nums, nums[:n]...)
	return nums[n:]
}

// 向右旋转 n 个字符
func rotateNRight(nums []int, n int) []int {
	// 1 2 3 4 5
	// 4 5 1 2 3

	// s1: 4 5 1 2 3 4 5
	// s2: 4 5 1 2 3
	n %= len(nums)
	nums = append(nums, nums[:n]...)
	return nums[n:]
}

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(rotateNLeft(s, 2))
}
