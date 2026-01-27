package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	// the LB port to connect to
	lbPort := os.Args[1]

	// write the message we want to deliver
	message := "hello my name is Amir"

	StartClient(message, lbPort)

}

func StartClient(message string, lbPort string) {

	// connect to the LB
	conn, err := net.Dial("tcp", ":"+lbPort)

	// check if the connection is correct
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Printf("connected to lb")

	defer conn.Close()

	// write into the connection
	conn.Write([]byte(message))
}
