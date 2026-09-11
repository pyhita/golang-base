package main

import (
	"context"
	helloworld2 "github.com/pyhita/golang-base/06_framework/grpc/greeter/helloworld"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50052"
	defaultName = "world"
)

func main() {
	// 建立到服务器的连接
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(), // 注意：WithInsecure 已废弃，建议使用 WithTransportCredentials
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	c := helloworld2.NewGreeterClient(conn)

	// 获取命令行参数或使用默认名称
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 设置带超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用远程方法
	r, err := c.SayHello(ctx, &helloworld2.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	log.Printf("响应结果: %s", r.Message)
}
