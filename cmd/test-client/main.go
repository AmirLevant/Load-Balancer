package main

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"
)

func main() {
	const lbPort = "8080"
	startClient(4, lbPort)

}

func startClient(data uint32, lbPort string) error {
	lbConn, err := net.Dial("tcp", ":"+lbPort)
	if err != nil {
		return err
	}
	defer lbConn.Close()

	slog.Info("Connected", slog.Any("address", lbConn.RemoteAddr()))

	txbuffer := make([]byte, 4)
	rxbuffer := make([]byte, 4)

	// Write the number 10 times to the lb
	for range 10 {
		// Serialise the number we want to send into bytes,
		// and send it over the connection
		binary.LittleEndian.PutUint32(txbuffer, data)

		_, err := lbConn.Write(txbuffer)
		if err != nil {
			return err
		}
	}

	// Indefinitely read from the lb
	for {
		_, err = lbConn.Read(rxbuffer)
		if err != nil {
			return err
		}

		msg := binary.LittleEndian.Uint32(rxbuffer)
		fmt.Printf("The msg from the server is: %d \n ", msg)
	}
}
