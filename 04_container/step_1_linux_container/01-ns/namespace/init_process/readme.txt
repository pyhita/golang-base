组合三个程序用来演示 init process 的特殊作用：

创建一个子进程并让其在隔离的 namespace 中执行新的命令：
ns_child_exec [options] command [arguments]
● options 指定隔离的 namespace
● command 指定子进程执行的命令

simple_init 程序实现了 init 的两个主要功能：
● 提供了一个简单的 shell 界面，用户可以通过该界面手动执行必要的 shell 命令来初始化命名空间
● 另一项功能是通过 waitpid() 来获取其子进程的状态信息。


演示这样一个现象：在 PID 命名空间中变得孤立的进程，会被 PID 命名空间中的 init 进程所接管，而不是被系统范围内的 init 进程所处理。

# ./ns_child_exec -p ./simple_init -v
    init: my PID is 1
init$ ./orphan
    init: created child 2
Parent (PID=2) created child with PID 3
Parent (PID=2; PPID=1) terminating
    init: SIGCHLD handler: PID 2 terminated
init$                   # simple_init prompt interleaved with output from child
Child  (PID=3) now an orphan (parent PID=1)
Child  (PID=3) terminating
    init: SIGCHLD handler: PID 3 terminated