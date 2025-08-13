package pkg1

import (
	"fmt"
	"github.com/pyhita/golang-base/01_golang_syntax/function/init_function/packag_init/pkg2"
)

var (
	_ = constInitCheck()
	v = variableInit("v")
)

const (
	c = "c"
)

func constInitCheck() string {
	if c != "" {
		fmt.Println("pkg1: const c init")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg1: var %s init\n", name)
	return name
}

func init() {
	pkg2.Add()
	fmt.Println("pkg1: init")
}
