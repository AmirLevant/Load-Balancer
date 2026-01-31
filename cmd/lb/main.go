package main

import (
	"lb/lb"
)

// sample packet for testing
type packet struct {
	length int
	data   []byte
}

func main() {
	var lbPort string = "8080"
	var serverPorts = []string{"9090", "9091", "9092"}
	lb.StartLoadBalancer(lbPort, serverPorts)
}
