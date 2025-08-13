package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// 练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。
//使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

var id = flag.String("id", "", "")

func main() {
	flag.Parse()

	var nodes []*html.Node
	for _, url := range flag.Args() {
		node, err := outline(url, *id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v", err)
			os.Exit(1)
		}
		nodes = append(nodes, node)
	}

	for _, node := range nodes {
		fmt.Println(node.Data)
		for _, a := range node.Attr {
			fmt.Println(a.Key, a.Val)
		}
	}
}

func outline(url, id string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	node := ElementByID(doc, id)
	if node := ElementByID(doc, id); node == nil {
		return nil, errors.New("node not found")
	}

	return node, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNodeV2(doc, id, findElement, findElement)
}

func forEachNodeV2(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if !pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if node := forEachNodeV2(c, id, pre, post); node != nil {
			return node
		}
	}

	if post != nil {
		if !post(n, id) {
			return n
		}
	}

	return nil
}

func findElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
	}
	return true
}
