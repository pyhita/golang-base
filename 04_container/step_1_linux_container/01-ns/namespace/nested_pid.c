/*
演示 PID Namespace 的嵌套关系：
在一个 PID 命名空间中，可以访问该命名空间内的所有进程，同时也可以访问那些属于子命名空间的进程。
这里的“访问”指的是能够执行对特定 PID 进行操作的系统调用（例如，使用 kill() 向某个进程发送信号）。
处于子 PID 命名空间的进程无法访问仅存在于父 PID 命名空间或更高层次命名空间的进程。

# ./multi_pidns 5
    # ls -d /proc4/[1-9]*        Topmost PID namespace created by program
    /proc4/1  /proc4/2  /proc4/3  /proc4/4  /proc4/5
    # ls -d /proc3/[1-9]*
    /proc3/1  /proc3/2  /proc3/3  /proc3/4
    # ls -d /proc2/[1-9]*
    /proc2/1  /proc2/2  /proc2/3
    # ls -d /proc1/[1-9]*
    /proc1/1  /proc1/2
    # ls -d /proc0/[1-9]*        Bottommost PID namespace
    /proc0/1

 清理资源：
 # umount /proc?
 # rmdir /proc?
*/

#define _GNU_SOURCE
#include <sched.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <string.h>
#include <signal.h>
#include <stdio.h>
#include <limits.h>
#include <sys/mount.h>
#include <sys/types.h>
#include <sys/stat.h>

/* A simple error-handling function: print an error message based
   on the value in 'errno' and terminate the calling process */

#define errExit(msg)    do { perror(msg); exit(EXIT_FAILURE); \
                        } while (0)

#define STACK_SIZE (1024 * 1024)

static char child_stack[STACK_SIZE];    /* Space for child's stack */
                /* Since each child gets a copy of virtual memory, this
                   buffer can be reused as each child creates its child */

/* Recursively create a series of child process in nested PID namespaces.
   'arg' is an integer that counts down to 0 during the recursion.
   When the counter reaches 0, recursion stops and the tail child
   executes the sleep(1) program. */

static int
childFunc(void *arg)
{
    static int first_call = 1;
    long level = (long) arg;

    if (!first_call) {

        /* Unless this is the first recursive call to childFunc()
           (i.e., we were invoked from main()), mount a procfs
           for the current PID namespace */

        char mount_point[PATH_MAX];

        snprintf(mount_point, PATH_MAX, "/proc%c", (char) ('0' + level));

        mkdir(mount_point, 0555);       /* Create directory for mount point */
        if (mount("proc", mount_point, "proc", 0, NULL) == -1)
            errExit("mount");
        printf("Mounting procfs at %s\n", mount_point);
    }

    first_call = 0;

    if (level > 0) {

        /* Recursively invoke childFunc() to create another child in a
           nested PID namespace */

        level--;
        pid_t child_pid;

        child_pid = clone(childFunc,
                    child_stack + STACK_SIZE,   /* Points to start of
                                                   downwardly growing stack */
                    CLONE_NEWPID | SIGCHLD, (void *) level);

        if (child_pid == -1)
            errExit("clone");

        if (waitpid(child_pid, NULL, 0) == -1)  /* Wait for child */
            errExit("waitpid");

    } else {

        /* Tail end of recursion: execute sleep(1) */

        printf("Final child sleeping\n");
        execlp("sleep", "sleep", "1000", (char *) NULL);
        errExit("execlp");
    }

    return 0;
}

int
main(int argc, char *argv[])
{
    long levels;

    levels = (argc > 1) ? atoi(argv[1]) : 5;
    childFunc((void *) levels);

    exit(EXIT_SUCCESS);
}
