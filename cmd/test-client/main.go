package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {

	// the LB port to connect to
	lbPort := "8080"

	// write the message we want to deliver
	message := 2

	StartClient(message, lbPort)

}

func StartClient(message int, lbPort string) {

	// connect to the LB
	conn, err := net.Dial("tcp", ":"+lbPort)

	// check if the connection is correct
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println("connected to lb")

	defer conn.Close()
	tempMessage := make([]byte, 1)

	for i := 0; i < 30; i++ {

		message = message + 3

		// convert int to message
		binary.LittleEndian.PutUint32(tempMessage, uint32(message))

		// write into the connection
		conn.Write(tempMessage)
	}
	finalMessage := make([]byte, 4)
	finalMessage = []byte("f")
	conn.Write(finalMessage)

}
