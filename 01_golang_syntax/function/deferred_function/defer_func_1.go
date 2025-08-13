package main

import (
	"fmt"
	"os"
	"sync"
)

// 示例一：在没有defer 机制的情况下，资源的回收会变得非常麻烦，容易出错
func writeToFile(fname string, data []byte, mu *sync.Mutex) error {
	mu.Lock()
	f, err := os.OpenFile(fname, os.O_RDWR, 0666)
	if err != nil {
		mu.Unlock()
		return err
	}

	_, err = f.Seek(0, 2)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	err = f.Sync()
	if err != nil {
		f.Close()
		mu.Unlock()
		return err
	}

	err = f.Close()
	if err != nil {
		mu.Unlock()
		return err
	}

	mu.Unlock()
	return nil
}

func main() {
	var mu sync.Mutex
	err := writeToFile("./foo.txt", []byte("hello, defer!\n"), &mu)
	if err != nil {
		fmt.Println("writeToFile error:", err)
		return
	}
	fmt.Println("writeToFile ok")
}
