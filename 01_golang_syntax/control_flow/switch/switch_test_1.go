package main

import "fmt"

func main() {
	n := 2
	switch n {
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 7:
		odd()
	case 2:
		fallthrough
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 8:
		even()
	default:
		unknown()
	}

	switch n {
	case 1, 3, 5, 7:
		odd()
	case 2, 4, 6, 8:
		even()
	default:
		unknown()
	}
}

func unknown() {

}

func even() {
	fmt.Println("even")
}

func odd() {
	fmt.Println("odd")
}
