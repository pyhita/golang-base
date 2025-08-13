package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

func main() {
	// for pprof
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	fmt.Printf("the number of goroutine: %d\n", runtime.NumGoroutine())

	ch := make(chan struct{})
	go func() {
		fmt.Println(mirroredQuery())
		// send exit signal
		ch <- struct{}{}
	}()

	<-ch
	fmt.Printf("the number of goroutine: %d\n", runtime.NumGoroutine())

	select {}
}

func mirroredQuery() string {
	responses := make(chan string)
	go func() { responses <- request("jd.com") }()
	go func() { responses <- request("taobao.com") }()
	go func() { responses <- request("baidu.com") }()
	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5)
	time.Sleep(time.Duration(n) * time.Second)

	//resp, err := http.Get("https://" + hostname)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//body, err := io.ReadAll(resp.Body)
	//return string(body[:10])
	return hostname
}
