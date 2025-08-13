package main

import (
	"errors"
	"fmt"
)

// error_01：构造 error 的方法

func main() {

	err1 := errors.New("new error")
	err2 := fmt.Errorf("new error: %w", err1)

	if errors.Is(err2, err1) {
		fmt.Println("error is err1")
	}

}
