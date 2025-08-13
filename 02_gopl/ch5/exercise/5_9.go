package main

import (
	"fmt"
	"os"
)

// 练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值。

func main() {
	fmt.Printf("result is %s\n", expand(os.Args[1], func(s string) string { return "xx" }))
}

func expand(s string, f func(string) string) string {
	return helper(s, 0, len(s), f)
}

// textfoohfoohh
func helper(s string, start, end int, f func(string) string) string {
	// len < 3 skip
	if end-start < 3 {
		return s
	}

	if s[start:start+3] == "foo" {
		t := s[start+3 : end]
		return f(s[start:start+3]) + helper(t, 0, len(t), f)
	}

	t := s[start+1 : end]
	return s[start:start+1] + helper(t, 0, len(t), f)
}
