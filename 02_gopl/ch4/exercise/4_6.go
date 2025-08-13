package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回

func main() {
	b := []byte("哈哈  哈 哈哈  a")
	b = replaceSpace(b)
	fmt.Printf("%s\n", b)
}

func replaceSpace(b []byte) []byte {
	for i := 0; i < len(b); {
		// 解码出字符
		first, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(first) {
			second, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(second) {
				// 覆盖第一个空格
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}

	return b
}
