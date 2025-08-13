package rotate

// 向右移动k位
// [1,2,3,4,5,6,7] => [5,6,7,1,2,3,4]
func rotate(nums []int, k int) {
	reverse(nums[len(nums)-k:])
	reverse(nums[:len(nums)-k])
	reverse(nums)
}

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
