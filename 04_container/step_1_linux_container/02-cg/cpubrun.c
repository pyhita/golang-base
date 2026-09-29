// 编译: gcc -O0 -pthread -o cpuburn cpuburn.c
// 用法: ./cpuburn 2      占用 2 个 CPU
//       ./cpuburn 1.5    占用 1.5 个 CPU
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <unistd.h>

#define PERIOD_NS 10000000L /* 10ms 一个周期 */

static long long now_ns(void) {
    struct timespec ts;
    clock_gettime(CLOCK_MONOTONIC, &ts);
    return (long long)ts.tv_sec * 1000000000LL + ts.tv_nsec;
}

/* 占满一个核 */
static void *spin(void *arg) {
    (void)arg;
    volatile unsigned long x = 0;
    for (;;) x++;
    return NULL;
}

/* 每个周期内忙循环 ratio 比例的时间,其余时间 sleep,用于占用小数核 */
static void *spin_ratio(void *arg) {
    double ratio = *(double *)arg;
    long long busy = (long long)(PERIOD_NS * ratio);
    struct timespec rest = {0, PERIOD_NS - busy};
    volatile unsigned long x = 0;
    for (;;) {
        long long start = now_ns();
        while (now_ns() - start < busy) x++;
        nanosleep(&rest, NULL);
    }
    return NULL;
}

int main(int argc, char **argv) {
    if (argc != 2) {
        fprintf(stderr, "usage: %s <n_cpus>\n", argv[0]);
        return 1;
    }
    double n = atof(argv[1]);
    if (n <= 0) {
        fprintf(stderr, "n 必须大于 0\n");
        return 1;
    }

    int full = (int)n;
    static double frac;
    frac = n - full;

    pthread_t tid;
    for (int i = 0; i < full; i++) {
        if (pthread_create(&tid, NULL, spin, NULL) != 0) {
            perror("pthread_create");
            return 1;
        }
    }
    if (frac > 1e-6) {
        if (pthread_create(&tid, NULL, spin_ratio, &frac) != 0) {
            perror("pthread_create");
            return 1;
        }
    }

    printf("pid=%d, 正在占用 %.2f 个 CPU, Ctrl+C 退出\n", getpid(), n);
    for (;;) pause();
}