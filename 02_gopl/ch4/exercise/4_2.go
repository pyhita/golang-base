package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

// 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法

type Algo string

const (
	SHA256 Algo = "sha256"
	SHA384 Algo = "sha384"
	SHA512 Algo = "sha512"
)

func main() {
	algo := flag.String("algo", "sha256", "specify the hash algorithm")
	flag.Parse()

	// receive input
	var s string
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		s = input.Text()
		fmt.Println("Input:", s)
	}

	switch Algo(*algo) {
	case SHA256:
		c1 := sha256.Sum256([]byte(s))
		fmt.Printf("%x\n", c1)
	case SHA384:
		c1 := sha512.Sum384([]byte(s))
		fmt.Printf("%x\n", c1)
	case SHA512:
		c1 := sha512.Sum512([]byte(s))
		fmt.Printf("%x\n", c1)
	default:
		fmt.Errorf("unknown hash algorithm: %s", *algo)
		os.Exit(1)
	}
}
