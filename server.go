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
	"github.com/hashicorp/go-uuid"
	pbgen "go-consul/pb_gen"
	"go-consul/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// 2.2 定义类
type Person struct {
	//继承服务端的代码
	pbgen.UnimplementedHelloServer
}

var cnt int

// 2.3绑定方法
func (Person) SayHello(ctx context.Context, person *pbgen.Person) (*pbgen.Person, error) {
	cnt += 1
	fmt.Printf("[%s]->调用第[%v]次\n", pt, cnt)
	person.Name += "hello"
	return person, nil
}

var pt string

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
	//获取空闲端口
	port, err := utils.GetFreePort()
	//3.写入注册服务信息
	srvId, _ := uuid.GenerateUUID()
	srvId = "test_gRPC" + srvId
	pt = fmt.Sprintf("%s:%v", "127.0.0.1", port)
	reg := api.AgentServiceRegistration{
		ID:      srvId,
		Name:    "test_gRPC_Consul",
		Tags:    []string{"test_gRPC", "test"},
		Port:    port,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			CheckID:  "consul_gRPC_test",
			GRPC:     fmt.Sprintf("%s:%v", "127.0.0.1", port),
			Timeout:  "1s",
			Interval: "5s",
		},
	}
	//4.注册服务到consul
	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		fmt.Println("注册服务到consul失败")
		return
	}

	//_____gRPC_______
	//1 初始化对象
	server := grpc.NewServer()
	//2.1注册服务
	pbgen.RegisterHelloServer(server, new(Person))
	fmt.Println(srvId + "运行")
	//3.设置监听，指定ip和port
	listen, err := net.Listen("tcp", ":"+fmt.Sprintf("%v", port))
	if err != nil {
		fmt.Println("监听出错")
		return
	}
	//健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	//4.启动服务
	go func() {
		err = server.Serve(listen)
		if err != nil {
			fmt.Println("启动服务失败")
			return
		}
	}()

	//退出时注销服务
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	err = consulClient.Agent().ServiceDeregister(srvId)
	if err != nil {
		fmt.Println("服务注销失败")
		return
	}
	fmt.Println("服务退出成功")

}
