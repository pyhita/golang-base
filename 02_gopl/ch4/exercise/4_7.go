package main

import (
	"fmt"
	"unicode/utf8"
)

// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？

// 额外申请内存版本
func reverseRuneV1(b []byte) []byte {
	z := make([]byte, len(b)) // 直接分配目标空间
	pos := 0
	for i := len(b); i > 0; {
		last, size := utf8.DecodeLastRune(b[:i])
		utf8.EncodeRune(z[pos:], last) // 直接写入目标位置
		i -= size
		pos += size
	}
	return z
}

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		// reverse byte
		reverseByte(b[i : i+size])
		i += size
	}

	reverseByte(b)
	return b
}

func reverseByte(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return b
}

func main() {
	s := "中国人"
	fmt.Println(string(reverseRuneV1([]byte(s))))

	t := "中国心"
	fmt.Println(string(reverseUTF8([]byte(t))))
}
