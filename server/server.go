package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {

	// we listen on port 8080
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server running on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	// always close the connection at the end
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err)
			}
			return
		}

		fmt.Printf("Message Recieved: %s", message)

	}

}
