package main

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"go-consul/setting"
)

func main() {
	config := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}
	clientCfg := constant.ClientConfig{
		NamespaceId:         "ab0ce7e9-1510-489f-91fe-cf1326400d68", // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": config,
		"clientConfig":  clientCfg,
	})
	if err != nil {
		fmt.Println("创建nacos对象失败")
		return
	}
	getConfig, err := client.GetConfig(vo.ConfigParam{
		DataId: "web-dev",
		Group:  "dev",
	})
	if err != nil {
		fmt.Println("获取配置失败")
		return
	}
	fmt.Println(getConfig)
	cfg := &setting.AppCfg{}
	json.Unmarshal([]byte(getConfig), &cfg)
	fmt.Println(cfg.Test)
	fmt.Println(cfg.Cur)
	//监控
	_ = client.ListenConfig(vo.ConfigParam{
		DataId: "web-dev",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("____________配置文件发生变化___________")
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
			json.Unmarshal([]byte(data), &cfg)
			fmt.Println(cfg.Test)
			fmt.Println(cfg.Cur)
		},
	})
	select {}
}
