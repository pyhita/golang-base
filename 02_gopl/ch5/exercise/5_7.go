package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

/*
练习 5.7： 完善startElement和endElement函数，
使其成为通用的HTML输出器。要求：输出注释结点，
文本结点以及每个元素的属性（< a href='...'>）。
使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。
编写测试，验证程序输出的格式正确。（详见11章）

<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <title>演示页面</title>
  <!-- 页面标题设置完成 -->
</head>
<body>

  <!-- 页面主内容开始 -->
  <h1 id="main-title" class="title">欢迎访问我的网站</h1>

  <p class="description">
    这是一个简单的 <strong>HTML 示例页面</strong>。
    <!-- strong 元素用于强调文本 -->
  </p>

  <a href="https://example.com" target="_blank" rel="noopener noreferrer">
    点击这里访问示例网站
  </a>

  <img src="https://example.com/logo.png" alt="示例图片" width="200" height="100">

  <!-- 页面底部 -->
  <footer>
    版权所有 © 2025
  </footer>

</body>
</html>


*/

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	outline2(doc)
}

func outline2(n *html.Node) {
	forEachNode(n, startElement, endElement)
}

func forEachNode(node *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		forEachNode(child, pre, post)
	}

	if post != nil {
		post(node)
	}
}

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		// build tag string
		var atr string
		for _, a := range n.Attr {
			atr = fmt.Sprintf("%s %s=%s", atr, a.Key, a.Val)
		}

		// 有孩子节点，才会输出
		if n.FirstChild != nil {
			fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, atr)
		} else {
			fmt.Printf("%*s<%s %s/>\n", depth*2, "", n.Data, atr)
		}
		depth++
	case html.TextNode:
		// text 节点，n.Data 就是text内容
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		depth++
	case html.CommentNode:
		depth++
		// <!-- strong 元素用于强调文本 -->
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		depth--
		// 有孩子节点，才会输出
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	case html.TextNode:
		depth--
		// no print
	case html.CommentNode:
		depth--
		// no print
	}
}
