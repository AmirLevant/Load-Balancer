package loadbalancer

import (
	"fmt"
	"log"
	"net"
)

func StartLoadBalancer(port string, serverPorts []string) {
	// Setting up Load Balancer
	loadBalancerListener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Printf("Error listening: %s", err)
	}

	defer loadBalancerListener.Close()
	fmt.Println("Load Balancer running on port:" + port)

	i := 0
	for {

		conn, err := loadBalancerListener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting: %s", err)
			continue
		}

		go HandleConnection(conn, serverPorts[i])

		if i == 2 {
			i = 0
		}
		i++
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

	fmt.Println("Load Balancer recieved Client Message")

	serverConn, err := net.Dial("tcp", ":"+serverPort)
	fmt.Println("LoadBalancer attempting to contact " + serverPort)
	if err != nil {
		fmt.Printf("Error Connecting: %s ", err)
	}

	defer serverConn.Close()

	serverConn.Write(RxBuffer)

}
