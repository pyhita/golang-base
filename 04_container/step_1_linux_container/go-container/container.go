package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cgV2()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/root/kante.yang/container/go-container/rootfs/"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func cgV2() {
	cgroups := "/sys/fs/cgroup"
	kante := filepath.Join(cgroups, "kante")

	// 创建新的 cgroup 目录
	must(os.MkdirAll(kante, 0755))

	// 启用 pids 控制器（只需在父目录设置一次）
	must(os.WriteFile(filepath.Join(cgroups, "cgroup.subtree_control"), []byte("+pids"), 0700))

	// 设置最大进程数
	must(os.WriteFile(filepath.Join(kante, "pids.max"), []byte("20"), 0700))

	// 设置自动释放（可选，v2 中 notify_on_release 已废弃，但仍可写入）
	_ = os.WriteFile(filepath.Join(kante, "notify_on_release"), []byte("1"), 0700)

	// 将当前进程加入该 cgroup
	must(os.WriteFile(filepath.Join(kante, "cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

