package main

import (
	"fmt"
	"log"
	"net/http"
)

// main函数将所有发送到/路径下的请求和handler函数关联起来
// 开头的请求其实就是所有发送到当前站点上的请求，服务监听8000端口
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
