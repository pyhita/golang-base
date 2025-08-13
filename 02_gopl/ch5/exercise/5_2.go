package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 练习 5.2： 编写函数，记录在HTML树中出现的同名元素的次数。

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	result := make(map[string]int)
	count(result, doc)

	for n, c := range result {
		fmt.Printf("%s: %d\n", n, c)
	}
}

func count(m map[string]int, node *html.Node) map[string]int {
	if node == nil {
		return m
	}

	if node.Type == html.ElementNode {
		m[node.Data]++
	}

	count(m, node.FirstChild)
	count(m, node.NextSibling)

	return m
}
