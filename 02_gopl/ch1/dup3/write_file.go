package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	data := `abc
abc
def
def
def
gh
gh`
	f, err := os.Create("text.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "open: %v\n", err)
	}
	defer f.Close()

	n, err := io.WriteString(f, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
	}

	fmt.Printf("wrote %d bytes\n", n)
}
