// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
package main

import "fmt"

func reverse(s *[8]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	reverse(&arr)
	fmt.Println(arr)
}
