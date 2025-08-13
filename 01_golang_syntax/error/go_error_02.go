package main

import "fmt"

// error_02：go 中 error 处理策略

func handleError01() {
	err := doSomething()
	if err != nil {
		// 不关心err变量底层错误值所携带的具体上下文信息
		// 执行简单错误处理逻辑并返回
		fmt.Println(err)
		return
	}
}

func handleError02() {
	err := doSomething()
	if err != nil {
		switch err.Error() {
		case "bufio: negative count":
			// ... ...
			return
		case "bufio: buffer full":
			// ... ...
			return
		case "bufio: invalid use of UnreadByte":
			// ... ...
			return
		default:
			// ... ...
			return
		}
	}
}

func handleError03() {
	
}

func doSomething() error {
	return nil
}

func main() {

}
