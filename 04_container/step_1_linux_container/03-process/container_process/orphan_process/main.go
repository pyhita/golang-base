package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// 父进程逻辑
	if os.Getenv("CHILD") == "" {
		runParent()
	} else {
		runChild()
	}
}

func runParent() {
	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Printf("im father pid=%d ppid=%d\n", pid, ppid)

	// 启动子进程
	cmd := exec.Command(os.Args[0])
	cmd.Env = append(os.Environ(), "CHILD=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to start child: %v\n", err)
		os.Exit(1)
	}

	// 父进程立即退出
	fmt.Println("father died..")
	os.Exit(0)
}

func runChild() {
	// 等待父进程退出
	time.Sleep(1 * time.Second)

	pid := os.Getpid()
	ppid := os.Getppid()
	fmt.Printf("im child pid=%d ppid=%d\n", pid, ppid)
}
