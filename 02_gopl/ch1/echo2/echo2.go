package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		// += 连接原字符串每次都会重新进行内存分配
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
