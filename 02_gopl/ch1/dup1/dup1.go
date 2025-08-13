package main

import (
	"bufio"
	"fmt"
	"os"
)

// 打印标准输入中多次出现的行，以重复次数开头
func main() {
	counter := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// 该变量从程序的标准输入中读取内容。每次调用 input.Scan()，即读入下一行，
	// 并移除行末的换行符；读取的内容可以调用 input.Text() 得到。
	// Scan 函数在读到一行时返回 true，不再有输入时返回 false。
	for input.Scan() {
		counter[input.Text()]++
		if input.Text() == "exit" {
			break
		}
	}

	for line, n := range counter {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
