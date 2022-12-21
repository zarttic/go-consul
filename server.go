/**
 *@filename       server.go
 *@Description		服务端
 *@author          liyajun
 *@create          2022-12-17 15:29:50
 */

package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	pb_gen "go-consul/pb_gen"
	"google.golang.org/grpc"
	"net"
)

// 2.2 定义类
type Person struct {
	//继承服务端的代码
	pb_gen.UnimplementedHelloServer
}

// 2.3绑定方法
func (Person) SayHello(ctx context.Context, person *pb_gen.Person) (*pb_gen.Person, error) {
	person.Name += "hello"
	return person, nil
}
func main() {
	fmt.Println("服务端 开启")
	//注册consul
	//1.初始化consul配置
	config := api.DefaultConfig()
	//创建consul对象
	consulClient, err := api.NewClient(config)
	if err != nil {
		fmt.Println("创建consul对象失败")
		return
	}
	//3.写入注册服务信息
	reg := api.AgentServiceRegistration{
		ID:      "test_gRPC",
		Name:    "test_gRPC_Consul",
		Tags:    []string{"test_gRPC"},
		Port:    8080,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			CheckID:  "consul_gRPC_test",
			TCP:      "127.0.0.1:8080",
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//4.注册服务到consul
	consulClient.Agent().ServiceRegister(&reg)

	//_____gRPC_______
	//1 初始化对象
	server := grpc.NewServer()
	//2.1注册服务
	pb_gen.RegisterHelloServer(server, new(Person))
	//3.设置监听，指定ip和port
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听出错")
		return
	}
	//4.启动服务
	err = server.Serve(listen)
	if err != nil {
		fmt.Println("启动服务失败")
		return
	}

}
