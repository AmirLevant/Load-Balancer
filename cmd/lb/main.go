package main

import (
	"lb/lb"
)

func main() {
	var lbPort string = "8080"
	var serverPorts = []string{"9090", "9091", "9092"}
	lb.StartLoadBalancer(lbPort, serverPorts)
}
