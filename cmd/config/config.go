package main

import (
	"fmt"
	"log"

	"lb/cmd/loadbalancer"
	"lb/cmd/server"

	"github.com/BurntSushi/toml"
)

type Config struct {
	LoadBalancer struct {
		Port int `toml:"port"`
	} `toml:"loadbalancer"`
	Servers []struct {
		Port int `toml:"port"`
	} `toml:"servers"`
}

func main() {

	// parsing config
	var cnfg Config
	if _, err := toml.DecodeFile("config.toml", &cnfg); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// setup LB port
	loadBalancerPort := fmt.Sprintf("%d", cnfg.LoadBalancer.Port)

	// Convert servers to string array
	serverPorts := make([]string, len(cnfg.Servers))
	for i, server := range cnfg.Servers {
		serverPorts[i] = fmt.Sprintf("%d", server.Port)
	}

	// run load balancer
	go loadbalancer.StartLoadBalancer(loadBalancerPort, serverPorts)

	for i := 0; i < len(serverPorts); i++ {
		go server.StartServer(serverPorts[i])
	}

	select {}
}
