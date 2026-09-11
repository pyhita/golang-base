#!/bin/bash
python /opt/test.py
 
# 添加这一样是为了防止上面那一行进程被干掉后，本进程还有代码在运行着，否则就也结束了
tail -f /dev/null 
