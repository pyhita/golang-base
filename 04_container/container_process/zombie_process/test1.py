#coding:utf-8
from multiprocessing import Process
import os
import time

def task(n):
    print("father--->%s son--->%s" %(os.getppid(),os.getpid()))
    time.sleep(n)

if __name__ == "__main__":
    # 首选必须知道一个知识点：一个python进程，就是一个python解释器进程，你的python代码本质都是字符串最终运行与调用的
    # 其实都是解释器的代码
    # 对于开启的子进程，即便你不去执行p.join()方法，也不用担心，因为只要python解释器还在运行着，它就会定期定期定期，
    # 注意是定期，去执行
    # 回收僵尸进程的功能，
    # 当如果你没有执行p.join()，同时也time.sleep(100000)住了主进程，那肯定就没有人来回收僵尸儿子了

    # 所以，如果你开启了一千个子进程，每个子进程都是运行一行打印功能然后在极短的时间内进入僵尸进程状态
    # 主进程在一个个开启这一千个子进程的过程中就会定期回收一部分僵尸儿子，当运行到最后执行到time.sleep(100000)时，主进程就停住了
    # 剩下的没来得及回收的僵尸儿子也就残留了，此处我想告诉大家的就是这件事，你不要以为你开一千个子进程每个都很快结束，然后最后你就
    # 会看到一千个僵尸儿子

    for i in range(1000):
        Process(target=task,args=(0,)).start()

    time.sleep(100000)