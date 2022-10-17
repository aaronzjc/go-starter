## 简介
`go-starter`是一款简单的Golang Web开发脚手架。基于Gin，集成多组件，目前涵盖

+ HTTP服务
+ GRPC服务
+ 日志组件
+ MySQL操作
+ HTTP请求

## 使用

安装构建器

```shell
# windows平台安装Task构建工具
go install github.com/go-task/task/v3/cmd/task@latest
# Linux平台使用原生Make构建工具
```

开发构建

```shell
# 构建api服务
task build:api
# 构建grpc服务
task build:grpc
# 生成grpc相关文件
task gen:grpc
```

运行
```
./bin/api.exe -c ./conf/dev.yml
./bin/grpc.exe -c ./conf/dev.yml
```