/**
 *@filename       consul-exit.go
 *@Description
 *@author          liyajun
 *@create          2022-12-18 16:28:26
 */

package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

// 注销consul服务
func main() {
	//1 初始化consul 配置
	config := api.DefaultConfig()
	//2 创建consul 对象
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("创建consul对象失败")
		return
	}
	//3 注销服务
	_ = client.Agent().ServiceDeregister("test_gRPC")
	fmt.Println("服务退出")
}
