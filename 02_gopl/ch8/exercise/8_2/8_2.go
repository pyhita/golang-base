package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Print(err)
		return
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		fmt.Printf("New connection from %s\n", conn.RemoteAddr())
	}

}
