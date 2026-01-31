package lb

import (
	"fmt"
	"log"
	"net"
)

func StartLoadBalancer(port string, serverPorts []string) {

	// Set up Load Balancer
	loadBalancerListener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Printf("Error listening: %s", err)
	}

	defer loadBalancerListener.Close()

	fmt.Println("Load Balancer running on port:" + port)

	// number dictating which server handles the request
	// increments with new connection made
	// increment means rotates server lb connects to

	serverTrackerNum := 0
	for {
		conn, err := loadBalancerListener.Accept()
		if err != nil {
			fmt.Printf("Error Accepting: %s", err)
			continue
		}

		fmt.Println("Load Balancer recieved a Client message")

		go HandleConnection(conn, serverPorts[serverTrackerNum])

		// reset the server cycle
		if serverTrackerNum == 2 {
			serverTrackerNum = 0
		}
		serverTrackerNum++
	}
}

func HandleConnection(clientConn net.Conn, serverPort string) {

	// close the connection at the end
	defer clientConn.Close()

	// connect to server
	serverConn, err := net.Dial("tcp", ":"+serverPort)
	fmt.Println("LoadBalancer attempting to contact server :" + serverPort)

	if err != nil {
		fmt.Printf("Error Connecting: %s ", err)
	}
	defer serverConn.Close()

	TxBuffer := make([]byte, 4)
	RxBuffer := make([]byte, 4)

	for {

		_, err := clientConn.Read(TxBuffer)

		if err != nil {
			log.Printf("Connection error: %v", err)
			return
		}

		serverConn.Write(TxBuffer)

		_, err = serverConn.Read(RxBuffer)

		if err != nil {
			log.Printf("error reading from server conn : %v", err)
			return
		}

		clientConn.Write(RxBuffer)

	}

}
