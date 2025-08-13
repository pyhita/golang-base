package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var (
	hostname   string
	serverName string
	port       int
	isOpen     bool
)

func main() {
	pflag.StringVarP(&hostname, "hostname", "h", "", "specific hostname")
	pflag.StringVarP(&serverName, "server-name", "s", "", "specific server name")
	pflag.IntVarP(&port, "port", "p", 8080, "specific port")
	pflag.BoolVarP(&isOpen, "open", "o", false, "open")
	pflag.Parse()

	fmt.Printf("%s - %s - %d - %t\n", hostname, serverName, port, isOpen)

}
