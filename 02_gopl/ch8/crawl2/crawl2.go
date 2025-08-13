package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pyhita/golang-base/02_gopl/ch5/links"
)

// 1 限制程序最多可以同时建立20个connect
// 2 确保程序在没有可爬取url的时候正常退出
func main() {
	var n int
	worklist := make(chan []string)

	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

// tokens
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	// acquire token
	tokens <- struct{}{}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	// release token
	<-tokens
	return list
}
