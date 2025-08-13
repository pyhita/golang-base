package main

import (
	"bufio"
	"fmt"
	"os"
)

type s struct {
	count int
	name  string
}

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
func main() {
	counter := make(map[string]*s)

	files := os.Args[1:]
	if len(files) == 0 {
		// read from stdin
		countLines(os.Stdin, counter, "os.stdin")
	} else {
		// read from files
		for _, f := range files {
			name := f
			f, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counter, name)
			f.Close()
		}
	}

	for line, n := range counter {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, n.name)
		}
	}
}

func countLines(f *os.File, counter map[string]*s, name string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		if counter[input.Text()] == nil {
			counter[input.Text()] = &s{count: 1, name: name}
		} else {
			counter[input.Text()].count++
		}
	}
}
