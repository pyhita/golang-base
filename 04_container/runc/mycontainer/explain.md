### 1 进入到mycontainer 目录
### 2 下载并解压 busybox rootfs 到该目录
### 3 执行 runc spec 生成 config.json，默认情况下就是到该目录下使用rootfs
### 4 执行 runc create mycontainer
### 5 执行 runc start mycontainer 启动容器