/* demo_userns.c

   Copyright 2013, Michael Kerrisk
   Licensed under GNU General Public License v2 or later

   Demonstrate the use of the clone() CLONE_NEWUSER flag.

   Link with "-lcap" and make sure that the "libcap-devel" (or
   similar) package is installed on the system.
*/

// 该程序会在新的用户命名空间中创建一个子账户。这个子账户只会显示其有效的用户和组 ID 以及其权限。
// 以普通用户身份运行这个程序时，会得到如下结果：
//
//     $ id -u          # Display effective user ID of shell process
//     1000
//     $ id -g          # Effective group ID of shell
//     1000
//     $ ./demo_userns
//     eUID = 65534;  eGID = 65534;  capabilities: =ep
#define _GNU_SOURCE
#include <sys/capability.h>
#include <sys/wait.h>
#include <sched.h>
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define errExit(msg)    do { perror(msg); exit(EXIT_FAILURE); \
                        } while (0)

static int                      /* Startup function for cloned child */
childFunc(void *arg)
{
    cap_t caps;

    for (;;) {
        printf("eUID = %ld;  eGID = %ld;  ",
                (long) geteuid(), (long) getegid());

        caps = cap_get_proc();
        printf("capabilities: %s\n", cap_to_text(caps, NULL));

        if (arg == NULL)
            break;

        sleep(5);
    }

    return 0;
}

#define STACK_SIZE (1024 * 1024)

static char child_stack[STACK_SIZE];    /* Space for child's stack */

int
main(int argc, char *argv[])
{
    pid_t pid;

    /* Create child; child commences execution in childFunc() */

    pid = clone(childFunc, child_stack + STACK_SIZE,    /* Assume stack
                                                           grows downward */
                CLONE_NEWUSER | SIGCHLD, argv[1]);
    if (pid == -1)
        errExit("clone");

    /* Parent falls through to here.  Wait for child. */

    if (waitpid(pid, NULL, 0) == -1)
        errExit("waitpid");

    exit(EXIT_SUCCESS);
}