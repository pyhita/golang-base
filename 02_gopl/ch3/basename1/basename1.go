package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	fmt.Println(s)
}
