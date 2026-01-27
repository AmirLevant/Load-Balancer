package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	var lbPort string = "8080"
	var serverPorts = []string{"9090", "9091", "9092"}
	StartLoadBalancer(lbPort, serverPorts)
}

func StartLoadBalancer(port string, serverPorts []string) {

	// Setting up Load Balancer
	loadBalancerListener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Printf("Error listening: %s", err)
	}

	defer loadBalancerListener.Close()
	fmt.Println("Load Balancer running on port:" + port)

	// number that dictates which server is handling the request
	// increments each new connection made
	// new increment means different server to connect to
	serverTrackerNum := 0
	for {
		conn, err := loadBalancerListener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting: %s", err)
			continue
		}

		go HandleConnection(conn, serverPorts[serverTrackerNum])

		// we reached the final server
		// reset to the first
		if serverTrackerNum == 2 {
			serverTrackerNum = 0
		}
		serverTrackerNum++
	}
}

func HandleConnection(clientConn net.Conn, serverPort string) {

	// always close the connection at the end
	defer clientConn.Close()

	RxBuffer := make([]byte, 1024)

	_, err := clientConn.Read(RxBuffer)

	if err != nil {
		log.Printf("Connection error: %v", err)
		return
	}

	fmt.Println("Load Balancer recieved a Client message")

	serverConn, err := net.Dial("tcp", ":"+serverPort)
	fmt.Println("LoadBalancer attempting to contact server :" + serverPort)

	if err != nil {
		fmt.Printf("Error Connecting: %s ", err)
	}

	defer serverConn.Close()

	serverConn.Write(RxBuffer)

}
