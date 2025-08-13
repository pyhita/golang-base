package main

func main() {

	var a int
	// 等同于 switch true
	switch {
	case a < 0:
		print("a < 0")
	case a == 0:
		print("a = 0")
	case a > 0:
		print("a > 0")
	default:
		print("invalid path")
	}

	switch i := 0; i + 2 {
	
	}
}
