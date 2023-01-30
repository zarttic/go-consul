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
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	pbgen "go-consul/pb_gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Println("客户端 开启")

		//_____gRPC_______
		//1.连接服务
		//grpc.WithInsecure() 已被弃用  使用如下代替
		//地址拼接
		dial, err := grpc.Dial(
			"consul://127.0.0.1:8500/test_gRPC_Consul?wait=14s&tag=test_gRPC",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), //负载均衡策略
		)
		if err != nil {
			fmt.Println("连接失败")
			return
		}
		//defer dial.Close()
		//2.初始化grpc客户端
		client := pbgen.NewHelloClient(dial)
		//3.调用远程函数
		p, err := client.SayHello(context.TODO(), &pbgen.Person{
			Name: "jxc",
			Age:  20,
		})
		if err != nil {
			fmt.Println("调用失败")
			return
		}
		fmt.Println(p)
		time.Sleep(time.Millisecond * 200)
	}
}
