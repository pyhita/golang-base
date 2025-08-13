package dedup

import (
	"bufio"
	"fmt"
	"os"
)

// map 模拟 set，不打印重复的行
func printLines() {
	seen := make(map[string]bool)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 读取一行输入
		text := input.Text()
		if !seen[text] {
			seen[text] = true
			fmt.Println(text)
		}
	}
}
