package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"log"
	_ "net/http/pprof" // Import pprof
	"os"
	"path"
	"sync"
	"time"
)

var verbose = pflag.BoolP("verbose", "v", false, "show verbose progress messages")

func main() {
	pflag.Parse()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	dirs := pflag.Args()
	fileSizes := make(chan int64)

	var wg sync.WaitGroup
	// producer
	for _, dir := range dirs {
		wg.Add(1)
		go walkDirs(dir, &wg, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nFiles, nBytes int64
loop:
	for {
		select {
		case <-done:
			// Drain all the goroutines, prevent goroutine leaks
			for range fileSizes {
				// do nothing
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nFiles++
			nBytes += size
		case <-tick:
			// print process
			printDiskUsage(nFiles, nBytes)
		}
	}
	printDiskUsage(nFiles, nBytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDirs(dir string, n *sync.WaitGroup, fileSizes chan int64) {
	defer n.Done()
	if canneled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := path.Join(dir, entry.Name())
			go walkDirs(subDir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return nil
	}
	return entries
}

var done = make(chan struct{})

func canneled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
