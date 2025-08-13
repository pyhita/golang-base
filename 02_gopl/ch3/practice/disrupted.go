package main

import "fmt"

// 练习 3.12： 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
func isDisrupted(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counter := make(map[rune]int)
	// build counter
	for _, c := range s {
		counter[c]++
	}

	for _, c := range t {
		counter[c]--
		if counter[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isDisrupted("a", "b"))
	fmt.Println(isDisrupted("abcd", "dcba"))
	fmt.Println(isDisrupted("abcd", "dcbad"))
}
