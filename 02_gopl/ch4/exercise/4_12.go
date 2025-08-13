package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

// 练习 4.12： 流行的web漫画服务xkcd也提供了JSON接口。例如，一个 https://xkcd.com/571/info.0.json
// 请求将返回一个很多人喜爱的571编号的详细描述。下载每个链接（只下载一次）然后创建一个离线索引。
// 编写一个xkcd工具，使用这些离线索引，打印和命令行输入的检索词相匹配的漫画的URL。

type XKCDComic struct {
	Month      string `json:"month"`      // 月份（字符串形式，如 "4"）
	Num        int    `json:"num"`        // 漫画编号（如 571）
	Link       string `json:"link"`       // 链接（示例中为空字符串）
	Year       string `json:"year"`       // 年份（字符串形式，如 "2009"）
	News       string `json:"news"`       // 新闻信息（示例中为空字符串）
	SafeTitle  string `json:"safe_title"` // 安全标题（不含特殊字符，如 "Can't Sleep"）
	Transcript string `json:"transcript"` // 漫画文本脚本（含多行和特殊符号）
	Alt        string `json:"alt"`        // 悬停提示文本（如关于 electric sheep 的提示）
	Img        string `json:"img"`        // 图片 URL（如 "https://imgs.xkcd.com/comics/cant_sleep.png"）
	Title      string `json:"title"`      // 完整标题（可能与 safe_title 相同）
	Day        string `json:"day"`        // 发布日期中的日（字符串形式，如 "20"）
	URL        string `json:"url"`        // 漫画的链接
}

const baseURL = "https://xkcd.com/"

// cache comic details
var details []*XKCDComic

func main() {
	start := flag.Int("start", 1, "search start index")
	end := flag.Int("end", 10, "search end index")
	search := flag.String("search", "", "search content")
	flag.Parse()

	// validate
	if *start > *end {
		fmt.Println("search start index is greater than end")
	}
	if *start <= 0 || *start > 1000 {
		fmt.Println("start must be between 0 and 1000")
	}
	if *end <= 0 || *end > 1000 {
		fmt.Println("end must be between 0 and 1000")
	}
	if *search == "" {
		fmt.Println("search is required")
	}

	buildLocalCache(*start, *end)
	searchComic(*search)
}

func buildLocalCache(start, end int) {
	var (
		wg      sync.WaitGroup
		results = make(chan *XKCDComic, end-start+1)
	)
	for i := start; i <= end; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("request %d comic\n", i)
			resp, err := http.Get(fmt.Sprintf("%s%d/info.0.json", baseURL, i))
			if err != nil {
				log.Printf("failed to fetch comic %d: %v", i, err)
				return
			}
			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read comic %d: %v", i, err)
				return
			}

			var x XKCDComic
			if err := json.Unmarshal(b, &x); err != nil {
				log.Printf("failed to parse comic %d: %v", i, err)
				return
			}

			x.URL = fmt.Sprintf("%s%d", baseURL, i)
			results <- &x
		}(i)
	}

	wg.Wait()
	close(results)

	for x := range results {
		details = append(details, x)
	}

	fmt.Printf("%d comics found\n", len(details))
}

func buildLocalCache2(start, end int) {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		results = make(chan *XKCDComic, end-start)
	)

	// 启动 goroutine 消费 results
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := start; i <= end; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			resp, err := http.Get(fmt.Sprintf("%s%d/info.0.json", baseURL, id))
			if err != nil {
				log.Printf("failed to fetch comic %d: %v", id, err)
				return
			}
			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read comic %d: %v", id, err)
				return
			}

			var x XKCDComic
			if err := json.Unmarshal(b, &x); err != nil {
				log.Printf("failed to parse comic %d: %v", id, err)
				return
			}

			x.URL = fmt.Sprintf("%s%d", baseURL, id)
			results <- &x
		}(i)
	}

	// 安全地收集结果
	for x := range results {
		mu.Lock()
		details = append(details, x)
		mu.Unlock()
	}

	fmt.Printf("%d comics found\n", len(details))
}

func searchComic(search string) {
	var found *XKCDComic
	for _, x := range details {
		if strings.Contains(strings.ToLower(x.Title), strings.ToLower(search)) {
			found = x
			break
		}

		if strings.Contains(strings.ToLower(x.Transcript), strings.ToLower(search)) {
			found = x
			break
		}

		if strings.Contains(strings.ToLower(x.SafeTitle), strings.ToLower(search)) {
			found = x
			break
		}
	}

	if found == nil {
		fmt.Println("no comic found")
		return
	}

	fmt.Printf("comic found url: %s\n", found.URL)
}
