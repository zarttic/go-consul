/**
 *@filename       client.go
 *@Description
 *@author          liyajun
 *@create          2022-12-17 15:36:02
 */

package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go-consul/pb_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
)

func main() {

	//consul
	//1.初始化consul配置
	config := api.DefaultConfig()
	//2.创建consul对象  此处可以修改一些配置
	consulClient, err := api.NewClient(config)
	if err != nil {
		fmt.Println("创建服务失败")
		return
	}
	//3.从consul上去获取健康的服务 [服务名称 str,别名 str,是否通过健康检查 bool,查询参数 QueryOptions]
	// [存储服务的切片（满足条件的服务端会有很多,可以添加简单的负载均衡）,额外查询的返回值，err]
	service, _, err := consulClient.Health().Service("test_gRPC_Consul", "test_gRPC", true, nil)
	if err != nil {
		fmt.Println("获取服务失败")
		return
	}

	//_____gRPC_______
	//1.连接服务
	//grpc.WithInsecure() 已被弃用  使用如下代替
	//地址拼接
	addr := service[0].Service.Address + ":" + strconv.Itoa(service[0].Service.Port)
	dial, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer dial.Close()
	//2.初始化grpc客户端
	client := pb_gen.NewHelloClient(dial)
	//3.调用远程函数
	p, err := client.SayHello(context.TODO(), &pb_gen.Person{
		Name: "jxc",
		Age:  20,
	})
	if err != nil {
		fmt.Println("调用失败")
		return
	}
	fmt.Println(p)
}
