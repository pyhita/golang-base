package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 利用 close() 广播退出信号
var verbose = flag.Bool("v", false, "show verbose progress messages")

var done = make(chan struct{})

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 监听用户的输入，确认是否退出
	go func() {
		// 阻塞接受用户输入
		os.Stdin.Read(make([]byte, 1))
		// 发送退出信号
		close(done)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(100 * time.Millisecond)
	}

	// Traverse the file tree.
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		// 如果tick是一个nil channel，相当于被忽略
		case <-tick:
			// 打印一次进度
			printDiskUsage(nfiles, nbytes)
		case size, ok := <-fileSizes:
			if !ok {
				// 说明channel已经关闭，停止等待
				break loop
			}
			nfiles++
			nbytes += size
		case <-done:
			// Drain fileSizes to allow existing goroutines to finish.
			// 还有很多goroutine阻塞在写入上，直接退出会造成泄漏
			for range fileSizes {
				// Do nothing.
			}
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	if canceled() {
		return
	}

	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var tokens = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	select {
	case tokens <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-tokens }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
