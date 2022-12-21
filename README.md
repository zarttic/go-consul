# go-consul
gRPC + consul测试 食用指南
> [gRPC官网](https://grpc.io/)
> 
> [Consul官网](https://www.consul.io/)
> 

> 
> 前置
安装依赖 
```shell
make init
```

> 编译proto文件

```shell
make proto
```

> 运行

1. 开启`consul`

```shell
consul agent -dev
```
2. 开启服务
```shell
make server
```

3. 开启客户端
```shell
make client
```
4. 注销服务
```shell
make exit
```
> 具体内容查看`Makefile`文件
