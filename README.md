# golang-base

Go 语言学习与实践代码仓库，内容覆盖基础语法、并发编程、`The Go Programming Language`（GOPL）练习、性能分析、容器、Kubernetes、Web 框架、IAM 以及 Webook 项目实践。

## 目录结构

- `01_golang_syntax`：Go 基础语法、类型、方法、接口、错误处理和并发示例
- `02_gopl`：GOPL 各章节示例与练习
- `03_pprof`：性能分析与 goroutine 相关示例
- `04_container`：容器原理及相关实验
- `05_k8s`：Kubernetes 资源配置与实践
- `06_framework`：Gin、GORM、gRPC 等框架示例
- `07_iam`：身份认证与授权相关示例
- `08_webook`：Webook 应用项目

## 环境要求

- Go 1.23 或更高版本
- 部分子项目包含独立的 `go.mod`，请在对应目录中执行命令

## 快速开始

克隆仓库并进入项目目录：

```bash
git clone git@github.com:pyhita/golang-base.git
cd golang-base
```

运行根目录的并发磁盘用量统计示例：

```bash
go run . [目录 ...]
```

使用 `-v` 参数可以定期输出处理进度：

```bash
go run . -v [目录 ...]
```

运行根模块测试：

```bash
go test ./...
```

## 说明

仓库中的示例主要用于学习和实验。不同目录可能有独立的依赖、运行方式或环境要求，请优先参考对应目录中的源码与配置文件。
