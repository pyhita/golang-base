package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pyhita/golang-base/02_gopl/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	// save pages
	if strings.Contains(url, "golang.org") {
		err = os.Mkdir("pages", 0755)
		if err != nil {
			log.Fatal(err)
		}

		for _, link := range list {
			resp, err := http.Get(link)
			if err != nil {
				log.Print(err)
				continue
			}
			defer resp.Body.Close()

			f, err := os.Create("pages/" + link)
			if err != nil {
				log.Print(err)
				continue
			}

			_, err = io.Copy(f, resp.Body)
			if err != nil {
				log.Print(err)
			}
		}
	}

	return list
}
