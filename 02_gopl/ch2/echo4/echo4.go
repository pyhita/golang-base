package main

import (
	"flag"
	"fmt"
	"strings"
)

// 调用flag.Bool函数会创建一个新的对应布尔型标志参数的变量。
// 它有三个属性：第一个是命令行标志参数的名字“n”，然后是该标志参数的默认值（这里是false），
// 最后是该标志参数对应的描述信息。如果用户在命令行输入了一个无效的标志参数，
// 或者输入-h或-help参数，那么将打印所有标志参数的名字、默认值和描述信息
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()

	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
