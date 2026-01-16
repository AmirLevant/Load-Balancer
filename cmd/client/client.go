package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// the LB port to connect to
	argPortClient := os.Args[1]

	// write the message we want to deliver
	message := "hello my name is Amir"

	// connect to the LB
	conn, err := net.Dial("tcp", ":"+argPortClient)

	// check if the connection is correct
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	defer conn.Close()

	// write into the connection
	conn.Write([]byte(message))

}
