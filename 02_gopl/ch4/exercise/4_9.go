package main

import (
	"bufio"
	"fmt"
	"os"
)

// 练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。
// 在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。

func main() {
	wordCount := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		text := input.Text()
		wordCount[text]++
	}

	for word, count := range wordCount {
		fmt.Printf("word: %s, count= %d\n", word, count)
	}
}
