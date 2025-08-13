package main

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprint(values) but adds commas.
func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intToString([]int{1, 2, 3}))
}
