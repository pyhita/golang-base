package main

import (
	"fmt"

	"golang.org/x/net/html"
)

// 练习5.12： gopl.io/ch5/outline2（5.5节）的startElement和
// endElement共用了全局变量depth，将它们修改为匿名函数，使其共享outline中的局部变量。

func outline3(node *html.Node) {
	var depth2 int

	startElement2 := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth2*2, "", n.Data)
			depth2++
		}
	}

	endElement2 := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth2--
			fmt.Printf("%*s</%s>\n", depth2*2, "", n.Data)
		}
	}

	forEachNode(node, startElement2, endElement2)
}
