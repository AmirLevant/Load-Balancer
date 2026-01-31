package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {

	// the LB port to connect to
	lbPort := "8080"

	// write the message we want to deliver
	var data uint32 = 7

	StartClient(data, lbPort)

}

func StartClient(data uint32, lbPort string) {

	// connect to the LB
	lbConn, err := net.Dial("tcp", ":"+lbPort)

	// check if the connection is correct
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	fmt.Println("connected to lb")

	defer lbConn.Close()

	Txbuffer := make([]byte, 4)
	Rxbuffer := make([]byte, 4)

	for i := 0; i < 10; i++ {

		// convert our data to bytes
		binary.LittleEndian.PutUint32(Txbuffer, data)

		// write into the connection
		lbConn.Write(Txbuffer)
	}

	for {
		lbConn.SetReadDeadline(time.Now().Add(time.Second))

		_, err = lbConn.Read(Rxbuffer)
		if err != nil {
			fmt.Println("Error connecting:", err)
			return
		}

		msg := binary.LittleEndian.Uint32(Rxbuffer)

		fmt.Printf("The msg from the server is: %d \n ", msg)

	}

}
