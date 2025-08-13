// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args 的第一个元素：os.Args[0]，是命令本身的名字；其它的元素则是程序启动时传给它的参数。
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
