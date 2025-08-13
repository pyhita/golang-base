// Boiling prints the boiling point of water.
package main

import "fmt"

/*
常量boilingF是在包一级范围声明语句声明的，
然后f和c两个变量是在main函数内部声明的声明语句声明的。
在包一级声明语句声明的名字可在整个包对应的每个源文件中访问，
而不是仅仅在其声明语句所在的源文件中访问。
相比之下，局部声明的名字就只能在函数内部很小的范围被访问。
*/
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}
