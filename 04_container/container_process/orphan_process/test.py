import os
import sys
import time

pid = os.getpid()
ppid = os.getppid()
print ('im father', 'pid', pid, 'ppid', ppid)
pid = os.fork()
#执行pid=os.fork()则会生成一个子进程
#返回值pid有两种值：
#    如果返回的pid值为0，表示在子进程当中
#    如果返回的pid值>0，表示在父进程当中
if pid > 0:
    print ('father died..')
    sys.exit(0)

# 保证主线程退出完毕
time.sleep(1)
print('im child', 'pid', os.getpid(), 'ppid', os.getppid())