package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "child" {
		child()
		return
	}

	parent()
}

func parent() {
	cmd := exec.Command("/proc/self/exe", "child")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER,

		// 将当前宿主机用户映射为 namespace 内的 root。
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},

		// 非 root 用户创建 user namespace 时通常需要禁止 setgroups。
		GidMappingsEnableSetgroups: false,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func child() {
	fmt.Printf("namespace 内 PID: %d\n", os.Getpid())
	fmt.Printf("UID: %d, GID: %d\n", os.Getuid(), os.Getgid())

	// 防止挂载操作传播到宿主机。
	if err := syscall.Mount(
		"",
		"/",
		"",
		syscall.MS_PRIVATE|syscall.MS_REC,
		"",
	); err != nil {
		log.Fatalf("设置 mount private 失败: %v", err)
	}

	if err := os.MkdirAll("/proc", 0555); err != nil {
		log.Fatalf("创建 /proc 失败: %v", err)
	}

	// 为新的 PID namespace 重新挂载 proc。
	if err := syscall.Mount(
		"proc",
		"/proc",
		"proc",
		0,
		"",
	); err != nil {
		log.Fatalf("挂载 /proc 失败: %v", err)
	}
	defer func() {
		if err := syscall.Unmount("/proc", 0); err != nil {
			log.Printf("卸载 /proc 失败: %v", err)
		}
	}()

	cmd := exec.Command("sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
