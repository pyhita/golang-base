#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#include <sys/capability.h>
#include <stdlib.h>
#include <fcntl.h>

#define STACK_SIZE (1024 * 1024)
static char child_stack[STACK_SIZE];

char* const child_args[] = {
    "/bin/bash",
    NULL
};

// 修正映射函数：在父进程中调用，使用正确的PID
void set_uid_map(pid_t pid, int inside_id, int outside_id, int length) {
    char path[256];
    sprintf(path, "/proc/%d/uid_map", pid);  // 使用传入的pid
    FILE* uid_map = fopen(path, "w");
    if (uid_map == NULL) {
        perror("fopen uid_map");
        return;
    }
    fprintf(uid_map, "%d %d %d", inside_id, outside_id, length);
    fclose(uid_map);
}

void set_gid_map(pid_t pid, int inside_id, int outside_id, int length) {
    char path[256];
    sprintf(path, "/proc/%d/gid_map", pid);  // 使用传入的pid
    FILE* gid_map = fopen(path, "w");
    if (gid_map == NULL) {
        perror("fopen gid_map");
        return;
    }
    fprintf(gid_map, "%d %d %d", inside_id, outside_id, length);
    fclose(gid_map);
}

int child_main(void* args) {
    printf("在子进程中!\n");
    sethostname("NewNamespace", 12);

    // 重新挂载/proc以确保PID命名空间隔离
    system("mount -t proc proc /proc");

    cap_t caps = cap_get_proc();
    printf("eUID = %ld; eGID = %ld; ", (long) geteuid(), (long) getegid());

    char *cap_text = cap_to_text(caps, NULL);
    printf("capabilities: %s\n", cap_text);
    cap_free(cap_text);
    cap_free(caps);

    execv(child_args[0], child_args);
    return 1;
}

int main() {
    printf("程序开始:\n");

    // 创建子进程
    pid_t child_pid = clone(child_main, child_stack + STACK_SIZE,
                          CLONE_NEWUSER | CLONE_NEWNS | CLONE_NEWPID |
                          CLONE_NEWIPC | CLONE_NEWUTS | SIGCHLD, NULL);

    if (child_pid == -1) {
        perror("clone");
        exit(1);
    }

    // 关键：在父进程中设置映射
    set_uid_map(child_pid, 0, getuid(), 1);
    set_gid_map(child_pid, 0, getgid(), 1);

    waitpid(child_pid, NULL, 0);
    printf("已退出\n");
    return 0;
}