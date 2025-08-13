package main

import (
	"bytes"
	"fmt"
)

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// 练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
func comma(s string) string {
	var buf bytes.Buffer

	for i := 0; i+3 < len(s); i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteByte(',')
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("12345"))
}
