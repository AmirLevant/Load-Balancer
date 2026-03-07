package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"net"
	"os"

	"github.com/BurntSushi/toml"
)

type serverConfig struct {
	ServerPort string `toml:"server_port"`
}

func main() {
	var cfg serverConfig
	if _, err := toml.DecodeFile("server.toml", &cfg); err != nil {
		slog.Error("Failed to decode server toml", slog.Any("error", err))
		os.Exit(1)
	}

	err := StartServer(cfg.ServerPort)
	if err != nil {
		slog.Error("Failed starting client", slog.Any("error", err))
		os.Exit(1)
	}
}

func StartServer(port string) error {
	// Listen on server port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed listening: %w", err)
	}
	defer listener.Close()

	slog.Info("Server running on port :" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Error Accepting:", slog.Any("error", err))
			continue
		}
		slog.Info("server " + port + " has recieved a connection")
		go func() {
			if err := handleConnection(conn); err != nil {
				slog.Error("Failed handling connection", slog.Any("error", err))
			}
		}()
	}
}

func handleConnection(conn net.Conn) error {
	// Always close the connection at the end
	defer func() {
		slog.Info("Closing connection")
		err := conn.Close()
		if err != nil {
			slog.Error("Failed closing connection:", slog.Any("error", err))
		} else {
			slog.Info("Connection closed")
		}
	}()

	// Len of data at the start of the buffer, 4 bytes
	rxBuffer := make([]byte, 4)
	txBuffer := make([]byte, 4)

	for i := 0; i < 10; i++ {
		// Read the message
		_, err := conn.Read(rxBuffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("Error reading: %w", err)
		}
		msg := binary.LittleEndian.Uint32(rxBuffer)
		slog.Info("Message from the client is: ", slog.Uint64("msg", uint64(msg)))

		// Write back
		msg = msg + 3
		binary.LittleEndian.PutUint32(txBuffer, msg)
		_, err = conn.Write(txBuffer)
		if err != nil {
			return fmt.Errorf("Error writing: %w", err)
		}

	}
	return nil

}
