## 简介
`go-starter`是一款简单的Golang Web开发脚手架。基于Gin，集成多组件，目前涵盖

+ 日志
+ MySQL
+ HTTP请求客户端

## 使用

安装构建器

```shell
go install github.com/go-task/task/v3/cmd/task@latest
```

构建

```shell
task build:dev
```

运行
```
./bin/api.exe -c ./conf/dev.yml
```