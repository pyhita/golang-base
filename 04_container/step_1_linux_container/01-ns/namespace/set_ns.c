// touch ~/uts
// mount --bind /proc/$$/ns/uts ~/uts

#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#include <fcntl.h>      // 需要添加这个头文件以使用O_RDONLY
#include <stdlib.h>     // 需要添加这个头文件以使用exit

int main(int argc, char *argv[])
{
    // 获取namespace文件描述符
    int fd = open(argv[1], O_RDONLY);
    if (fd == -1) {
        perror("open");
        exit(EXIT_FAILURE);
    }

    // 加入新的namespace
    if (setns(fd, 0) == -1) {
        perror("setns");
        close(fd);
        exit(EXIT_FAILURE);
    }
    close(fd);

    execvp(argv[2], &argv[2]);

    // 如果execvp返回，说明执行失败
    perror("execvp");
    exit(EXIT_FAILURE);
}


// ./setns ~/uts /bin/bash