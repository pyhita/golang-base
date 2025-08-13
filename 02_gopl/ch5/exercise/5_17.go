package main

import "golang.org/x/net/html"

// 练习5.17： 编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，
// 返回与这些标签名匹配的所有元素。下面给出了2个例子：

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	nodes := []*html.Node{}
	if doc == nil {
		return nodes
	}

	if doc.Type == html.ElementNode {
		if contains(doc.Data, name) {
			nodes = append(nodes, doc)
		}
	}

	nodes = append(nodes, ElementsByTagName(doc, name...)...)
	nodes = append(nodes, ElementsByTagName(doc, name...)...)

	return nodes
}

func contains(target string, strs []string) bool {
	for _, str := range strs {
		if target == str {
			return true
		}
	}
	return false
}
