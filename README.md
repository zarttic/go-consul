# go-consul
gRPC + consul测试 食用指南
> gRPC官网
https://grpc.io/

> 前置
安装依赖 
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

> 编译proto文件

```shell
protoc --go_out=. --go-grpc_out=. *.proto
```

> 运行

开启`consul`
```shell
go run server.go
go run client.go
```
