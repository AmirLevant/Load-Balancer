package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	serverPort := os.Args[1]

	StartServer(serverPort)
}

func StartServer(port string) {

	// we listen on port 8080
	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server running on port :" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err)
			continue
		}
		fmt.Println("server " + port + " has recieved a message")
		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) {
	// always close the connection at the end
	defer conn.Close()

	message := make([]byte, 1)

	for {

		_, err := conn.Read(message)

		if err != nil && err != io.EOF {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Printf("Message Content is: %s \n", message)
	}

}
