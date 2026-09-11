#coding:utf-8
from multiprocessing import Process
import os
import time
 
def task(n):
    print("father--->%s son--->%s" %(os.getppid(),os.getpid()))
    time.sleep(n)
 
if __name__ == "__main__":
    p1=Process(target=task,args=(10,))
    p2=Process(target=task,args=(10,))
    p3=Process(target=task,args=(10,))
    p1.start()
    p2.start()
    p3.start()
    print("main--->%s" %os.getpid())
    time.sleep(10000)
