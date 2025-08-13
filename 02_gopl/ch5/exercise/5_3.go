package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

// 练习 5.3： 编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素，因为这些元素对浏览者是不可见的。

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, text := range visitText(nil, doc) {
		fmt.Println(text)
	}
}

func visitText(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}

	if n.Type == html.TextNode {
		if n.Data != "script" && n.Data != "style" {
			for _, attr := range n.Attr {
				texts = append(texts, attr.Val)
			}
		}
	}

	texts = visitText(texts, n.FirstChild)
	texts = visitText(texts, n.NextSibling)

	return texts
}
