package main

import "fmt"

// 练习5.16： 编写多参数版本的strings.Join。

func main() {
	fmt.Println(join("-", 1, 2, "xxx"))
}

func join(sep string, strs ...interface{}) string {
	var res string
	for i, str := range strs {
		if i == 0 {
			res = fmt.Sprintf("%v", str)
			continue
		}
		res = fmt.Sprintf("%s%s%v", res, sep, str)
	}

	return res
}
