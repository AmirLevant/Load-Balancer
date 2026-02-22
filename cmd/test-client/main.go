package main

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/BurntSushi/toml"
)

type clientConfig struct {
	Lb_address string `toml:"lb_address"`
}

func main() {

	time.Sleep(2 * time.Second)
	var cfg clientConfig

	if _, err := toml.DecodeFile("client.toml", &cfg); err != nil {
		slog.Error("failed to decode client toml", slog.Any("error", err))
	}
	err := startClient(4, cfg.Lb_address)
	if err != nil {
		slog.Error("Failed starting client", slog.Any("error", err))
	}

}

func startClient(data uint32, lbAddress string) error {
	lbConn, err := net.Dial("tcp", lbAddress)
	if err != nil {
		slog.Error("Failed dialing lb", slog.Any("error", err))
		return err
	}
	defer func() {
		fmt.Println("Closing connection")
		err := lbConn.Close()
		if err != nil {
			fmt.Println("Failed closing connection:", err)
		} else {
			fmt.Println("Connection closed")
		}
	}()

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
			return fmt.Errorf("writing to lb: %w", err)
		}
	}

	// read from the lb
	for range 10 {
		_, err = lbConn.Read(rxbuffer)
		if err != nil {
			return fmt.Errorf("reading from lb: %w", err)
		}

		msg := binary.LittleEndian.Uint32(rxbuffer)
		fmt.Printf("The msg from the server is: %d \n ", msg)
	}
	return nil
}
