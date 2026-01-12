package main

import (
	"fmt"
	"log"
	"net"

	"github.com/burntsushi/toml"
)

type Config struct {
	loadBalancer struct {
		Port int `toml: "port"`
	} `toml:"loadbalancer"`
	Servers []struct {
		Port int `toml: "port"`
	} `toml:"servers"`
}

func main() {

	var cnfg Config
	if _, err := toml.DecodeFile("config.toml", &cnfg); err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// take load balancer port from config, 8080
	// listen on port
	loadBalancerPort := fmt.Sprintf(":%d", cnfg.loadBalancer.Port)
	ClientListener, err := net.Listen("tcp", loadBalancerPort)

	if err != nil {
		fmt.Printf("Error listening: %s", err)
	}

	defer ClientListener.Close()

	fmt.Println("Load Balancer running on port :8080")

	for {
		conn, err := ClientListener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting: %s", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// always close the connection at the end
	defer conn.Close()

	remoteAdd := conn.RemoteAddr()
	localAdd := conn.LocalAddr()

	fmt.Printf("The remote address in the client is : %s\n", remoteAdd)
	fmt.Printf("The local address in the client is: %s\n", localAdd)

	RxBuffer := make([]byte, 1024)

	_, err := conn.Read(RxBuffer)

	if err != nil {
		log.Printf("Connection error: %v", err)
		return
	}

	fmt.Printf("I got from the client: %s", RxBuffer)

	conn1, err := net.Dial("tcp", ":9090")

	if err != nil {
		fmt.Printf("Error Connecting: %s ", err)
	}

	defer conn1.Close()

	conn1.Write(RxBuffer)

}
