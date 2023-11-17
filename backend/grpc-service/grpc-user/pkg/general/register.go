package general

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

func Register(serverIP string, serverPort int, serverName string, serverTags []string, serverID string, consulConnStr string) {
	cfg := api.DefaultConfig()
	cfg.Address = consulConnStr // 这里是Consul的IP和地址
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", serverIP, serverPort), // 服务端的健康检查 支持HTTP
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "3s",
	}
	// 生成注册对象
	reg := new(api.AgentServiceRegistration)
	reg.Address = serverIP
	reg.ID = serverID
	reg.Port = serverPort
	reg.Name = serverName
	reg.Tags = serverTags
	reg.Check = check

	err = client.Agent().ServiceRegister(reg)
	if err != nil {
		panic(err)
	}
}

func DeregisterService(srvID string, consulConnStr string) error {
	config := api.DefaultConfig()
	config.Address = consulConnStr // 这里是Consul的IP和地址
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agent := client.Agent()

	err = agent.ServiceDeregister(srvID)
	if err != nil {
		return err
	}

	fmt.Println("Service deregistered successfully!")
	return nil
}
