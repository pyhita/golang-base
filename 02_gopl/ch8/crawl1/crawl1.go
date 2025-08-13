package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pyhita/golang-base/02_gopl/ch5/links"
)

// 并发的爬虫
func main() {
	worklist := make(chan []string)
	go func() { worklist <- os.Args[1:] }()

	depth := 3
	l := len(worklist)
	seen := make(map[string]bool)

	for i := 0; i < depth; i++ {
		for j := 0; j < l; j++ {
			list := <-worklist
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
		}
		l = len(worklist)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
