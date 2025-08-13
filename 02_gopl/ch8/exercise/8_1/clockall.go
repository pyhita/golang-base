package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/*
$ TZ=US/Eastern    ./clock2 -port 8010 &
$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
*/
type server struct {
	name string
	addr string
	msg  string
}

func main() {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	servers := parse(args)
	// conncet and read
	for _, srv := range servers {
		conn, _ := net.Dial("tcp", srv.addr)
		defer conn.Close()
		go func(s *server) {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				s.msg = scanner.Text()
			}
		}(srv)
	}

	for {
		fmt.Println()
		for _, srv := range servers {
			fmt.Printf("%s: %s\n", srv.name, srv.msg)
		}
		fmt.Printf("-------------")
		time.Sleep(1 * time.Second)
	}
}

func parse(args []string) []*server {
	var servers []*server
	for _, arg := range args {
		s := strings.Split(arg, "=")
		servers = append(servers, &server{
			name: s[0],
			addr: s[1],
		})
	}
	return servers
}
