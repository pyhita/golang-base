package main

import "fmt"

// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。

func main() {
	s := []string{"a", "b", "b", "b", "c", "b"}
	s = remove(s)
	fmt.Println(s)
}

func remove(s []string) []string {
	for i := 0; i < len(s)-1; {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
		} else {
			i++
		}
	}

	return s
}
