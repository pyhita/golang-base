package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const numItems = 100000 // 控制数据量，防止内存溢出
	data := make([]string, numItems)
	for i := range data {
		data[i] = "xxxxxxxxxx"
	}

	// 使用 `+=` 进行字符串拼接
	t1 := time.Now()
	var res1, sep string
	for _, d := range data {
		res1 += sep + d
		sep = "--"
	}
	t2 := time.Since(t1)

	// 使用 `strings.Join` 进行拼接
	t3 := time.Now()
	res2 := strings.Join(data, "--")
	t4 := time.Since(t3)

	fmt.Printf("Using '+=' concatenation: %v\n", t2)
	fmt.Printf("Using 'strings.Join': %v\n", t4)
	fmt.Printf("Lengths: res1=%d, res2=%d\n", len(res1), len(res2))
}
