package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()

	for i := 0; i < len(pflag.Args()); i++ {
		fmt.Println(pflag.Arg(i))
	}
}
