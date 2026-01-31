package main

import (
	"encoding/binary"
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

	// len of data at the start of the buffer, 4 bytes
	Rxbuffer := make([]byte, 4)
	Txbuffer := make([]byte, 4)

	for i := 0; i < 10; i++ {

		// read the message
		_, err := conn.Read(Rxbuffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading:", err)
			return
		}
		msg := binary.LittleEndian.Uint32(Rxbuffer)
		fmt.Printf("Message Content is: %d \n", msg)

		// write back
		msg = msg + 3
		binary.LittleEndian.PutUint32(Txbuffer, msg)
		conn.Write(Txbuffer)
	}

}
