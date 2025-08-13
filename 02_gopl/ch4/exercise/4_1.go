package main

import (
	"crypto/sha256"
	"fmt"
)

// 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)

func countDiffBits(s, t string) int {
	c1 := sha256.Sum256([]byte(s))
	c2 := sha256.Sum256([]byte(t))

	var count int
	for i := 0; i < len(c1); i++ {
		b1 := c1[i]
		b2 := c2[i]

		for j := 0; j < 8; j++ {
			c1 := b1 & (1 << uint(j))
			c2 := b2 & (1 << uint(j))

			if c1 != c2 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countDiffBits("xxxx", "xxxx"))
	fmt.Println(countDiffBits("x", "X"))
}
