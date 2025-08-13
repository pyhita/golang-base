package main

import "fmt"

// 删除slice 中间位置的某个元素，并保持原来的顺序
func removeOrderMid(nums []int, k int) []int {
	copy(nums[k:], nums[k+1:])

	return nums[:len(nums)-1]
}

// 删除slice 中间位置的某个元素，不用保持既定的顺序
// 直接用最后一个元素覆盖之前的元素
func removeNoneOrderMid(nums []int, k int) []int {
	nums[k] = nums[len(nums)-1]

	return nums[:len(nums)-1]
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(removeOrderMid(s, 2)) // "[5 6 8 9]"

	s = []int{5, 6, 7, 8, 9}
	fmt.Println(removeNoneOrderMid(s, 2)) // "[5 6 9 8]"
}
