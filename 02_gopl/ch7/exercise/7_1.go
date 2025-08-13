package main

import (
	"bufio"
	"bytes"
	"fmt"
)

// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*c++
	}
	return len(p), nil
}

func main() {
	s := "Hello, World!\nHello, 世界！"

	var wc WordCounter
	fmt.Fprintf(&wc, s)
	fmt.Println(wc)

	var lc LineCounter
	fmt.Fprintf(&lc, s)
	fmt.Println(lc)
}
